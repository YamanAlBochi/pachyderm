package dbutil

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pachyderm/pachyderm/v2/src/internal/backoff"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
)

var (
	txStartedMetric = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "pachyderm",
		Subsystem: "postgres",
		Name:      "tx_start_count",
		Help:      "Count of transactions that have been started.  One transaction may start many underlying database transactions, whose status is tracked separately.",
	})
	txFinishedMetric = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "pachyderm",
		Subsystem: "postgres",
		Name:      "tx_finish_count",
		Help:      "Count of transactions that have finished, by outcome ('error', 'ok', etc.).",
	}, []string{"outcome"})
	txDurationMetric = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "pachyderm",
		Subsystem: "postgres",
		Name:      "tx_duration_seconds",
		Help:      "Time taken for a transaction, by outcome ('error', 'ok', etc.).",
		Buckets: []float64{
			0.0001, // 100us
			0.0005, // .5ms
			0.001,  // 1ms
			0.002,  // 2ms
			0.005,  // 5ms
			0.01,   // 10ms
			0.02,   // 20ms
			0.05,   // 50ms
			0.1,    // 100ms
			0.2,    // 200ms
			0.5,    // 500ms
			1,      // 1s
			2,      // 2s
			5,      // 5s
			30,     // 30s
			60,     // 60s
			300,    // 5m
			600,    // 10m
			3600,   // 1h
			86400,  // 1d
		},
	}, []string{"outcome"})

	underlyingTxStartedMetric = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "pachyderm",
		Subsystem: "postgres",
		Name:      "tx_underlying_start_count",
		Help:      "Count of underlying database transactions that have been started.",
	})
	underlyingTxFinishMetric = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "pachyderm",
		Subsystem: "postgres",
		Name:      "tx_underlying_finish_count",
		Help:      "Count of underlying database transactions that have finished, by outcome ('commit_ok', 'commit_failed', 'rollback_ok', 'rollback_failed', 'failed_start', etc.)",
	}, []string{"outcome"})
)

type withTxConfig struct {
	sql.TxOptions
	backoff.BackOff
}

// WithTxOption parameterizes the WithTx function
type WithTxOption func(c *withTxConfig)

// WithIsolationLevel runs the transaction with the specified isolation level.
func WithIsolationLevel(x sql.IsolationLevel) WithTxOption {
	return func(c *withTxConfig) {
		c.TxOptions.Isolation = x
	}
}

// WithReadOnly causes WithTx to run the transaction as read only
func WithReadOnly() WithTxOption {
	return func(c *withTxConfig) {
		c.TxOptions.ReadOnly = true
	}
}

// WithBackOff sets the BackOff used when retrying
func WithBackOff(bo backoff.BackOff) WithTxOption {
	return func(c *withTxConfig) {
		c.BackOff = bo
	}
}

// WithTx calls cb with a transaction,
// The transaction is committed IFF cb returns nil.
// If cb returns an error the transaction is rolled back.
func WithTx(ctx context.Context, db *sqlx.DB, cb func(tx *sqlx.Tx) error, opts ...WithTxOption) error {
	backoffStrategy := backoff.NewExponentialBackOff()
	backoffStrategy.InitialInterval = 10 * time.Millisecond
	backoffStrategy.MaxElapsedTime = 0
	c := &withTxConfig{
		TxOptions: sql.TxOptions{
			Isolation: sql.LevelSerializable,
		},
		BackOff: backoffStrategy,
	}
	for _, opt := range opts {
		opt(c)
	}
	start := time.Now()
	txStartedMetric.Inc()
	err := backoff.RetryUntilCancel(ctx, func() error {
		underlyingTxStartedMetric.Inc()
		tx, err := db.BeginTxx(ctx, &c.TxOptions)
		if err != nil {
			underlyingTxFinishMetric.WithLabelValues("failed_start").Inc()
			return err
		}
		return tryTxFunc(tx, cb)
	}, c.BackOff, func(err error, _ time.Duration) error {
		if isTransactionError(err) {
			return nil
		}
		return err
	})
	duration := time.Since(start).Seconds()
	if err != nil {
		// Inspecting err could yield a better outcome type than "error", but some care is
		// needed.  For example, `cb` could return "context deadline exceeded" because it
		// creates a sub-context that expires, and that's a different error than 'commit'
		// failing because the deadline expired during commit.
		txFinishedMetric.WithLabelValues("error").Inc()
		txDurationMetric.WithLabelValues("error").Observe(duration)
		return err
	}
	txFinishedMetric.WithLabelValues("ok").Inc()
	txDurationMetric.WithLabelValues("ok").Observe(duration)
	return nil
}

func tryTxFunc(tx *sqlx.Tx, cb func(tx *sqlx.Tx) error) error {
	if err := cb(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			underlyingTxFinishMetric.WithLabelValues("rollback_failed").Inc()
			logrus.Error(rbErr)
			return err // The user error, not the rollback error.
		}
		underlyingTxFinishMetric.WithLabelValues("rollback_ok").Inc()
		return err
	}
	if err := tx.Commit(); err != nil {
		underlyingTxFinishMetric.WithLabelValues("commit_failed").Inc()
		return err
	}
	underlyingTxFinishMetric.WithLabelValues("commit_ok").Inc()
	return nil
}

func isTransactionError(err error) bool {
	pqerr := &pq.Error{}
	if errors.As(err, &pqerr) {
		return pqerr.Code.Class() == "40"
	}
	return IsErrTransactionConflict(err)
}

// ErrTransactionConflict should be used by user code to indicate a conflict in
// the transaction that should be reattempted.
type ErrTransactionConflict struct{}

func (err ErrTransactionConflict) Is(other error) bool {
	_, ok := other.(ErrTransactionConflict)
	return ok
}

func (err ErrTransactionConflict) Error() string {
	return "transaction conflict, will be reattempted"
}

// IsErrTransactionConflict determines if an error is an ErrTransactionConflict error
func IsErrTransactionConflict(err error) bool {
	return errors.Is(err, ErrTransactionConflict{})
}
