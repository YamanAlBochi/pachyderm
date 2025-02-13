package v2_6_0

import (
	"context"

	"github.com/pachyderm/pachyderm/v2/src/pps"

	"github.com/pachyderm/pachyderm/v2/src/internal/errors"
	"github.com/pachyderm/pachyderm/v2/src/internal/pachsql"
)

func branchlessCommitsPPS(ctx context.Context, tx *pachsql.Tx) error {
	jis, err := listCollectionProtos(ctx, tx, "jobs", &pps.JobInfo{})
	if err != nil {
		return errors.Wrap(err, "collecting jobs")
	}
	for _, ji := range jis {
		// TODO(provenance): nil commit.Branch field in storage
		ji.OutputCommit.Repo = ji.OutputCommit.Branch.Repo
		if err := updateCollectionProto(ctx, tx, "jobs", jobKey(ji.Job), jobKey(ji.Job), ji); err != nil {
			return errors.Wrapf(err, "update collections.jobs with key %q", jobKey(ji.Job))
		}
	}
	pis, err := listCollectionProtos(ctx, tx, "pipelines", &pps.PipelineInfo{})
	if err != nil {
		return errors.Wrap(err, "collecting pipelines")
	}
	for _, pi := range pis {
		pi.SpecCommit.Repo = pi.SpecCommit.Branch.Repo
		if err := updateCollectionProto(ctx, tx, "pipelines", pipelineCommitKey(pi.SpecCommit), pipelineCommitKey(pi.SpecCommit), pi); err != nil {
			return errors.Wrapf(err, "update collections.pipelines with key %q", pipelineCommitKey(pi.SpecCommit))
		}
	}
	return nil
}
