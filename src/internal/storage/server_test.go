package storage

import (
	"strings"
	"testing"

	"github.com/pachyderm/pachyderm/v2/src/internal/dockertestenv"
	"github.com/pachyderm/pachyderm/v2/src/internal/pachconfig"
	"github.com/pachyderm/pachyderm/v2/src/internal/pctx"
	"github.com/pachyderm/pachyderm/v2/src/internal/require"
	"github.com/pachyderm/pachyderm/v2/src/internal/storage/fileset"
	"github.com/pachyderm/pachyderm/v2/src/internal/storage/track"
)

func TestServer(t *testing.T) {
	ctx := pctx.TestContext(t)
	db := dockertestenv.NewTestDB(t)
	objC := dockertestenv.NewTestObjClient(ctx, t)

	tracker := track.NewTestTracker(t, db)
	fileset.NewTestStorage(ctx, t, db, tracker)

	var config pachconfig.StorageConfiguration
	s, err := New(Env{
		DB:          db,
		ObjectStore: objC,
	}, config)
	require.NoError(t, err)

	w := s.Filesets.NewWriter(ctx)
	require.NoError(t, w.Add("test.txt", "", strings.NewReader("hello world")))
	id, err := w.Close()
	require.NoError(t, err)
	t.Log(id)
}
