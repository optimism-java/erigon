package health

import (
	"context"
	"github.com/ledgerwatch/erigon-lib/common/hexutil"

	"github.com/optimism-java/erigon/rpc"
)

type NetAPI interface {
	PeerCount(_ context.Context) (hexutil.Uint, error)
}

type EthAPI interface {
	GetBlockByNumber(_ context.Context, number rpc.BlockNumber, fullTx bool) (map[string]interface{}, error)
	Syncing(ctx context.Context) (interface{}, error)
}
