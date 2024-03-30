package bodydownload_test

import (
	"testing"

	"github.com/optimism-java/erigon/turbo/stages/bodydownload"
	"github.com/optimism-java/erigon/turbo/stages/mock"
	"github.com/stretchr/testify/require"

	"github.com/optimism-java/erigon/consensus/ethash"
)

func TestCreateBodyDownload(t *testing.T) {
	t.Parallel()
	m := mock.Mock(t)
	tx, err := m.DB.BeginRo(m.Ctx)
	require.NoError(t, err)
	defer tx.Rollback()
	bd := bodydownload.NewBodyDownload(ethash.NewFaker(), 128, 100, m.BlockReader, m.Log)
	if _, _, _, _, err := bd.UpdateFromDb(tx); err != nil {
		t.Fatalf("update from db: %v", err)
	}
}
