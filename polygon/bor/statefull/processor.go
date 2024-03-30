package statefull

import (
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/optimism-java/erigon/consensus"
	"github.com/optimism-java/erigon/core/types"
)

type ChainContext struct {
	Chain consensus.ChainReader
	Bor   consensus.Engine
}

func (c ChainContext) Engine() consensus.Engine {
	return c.Bor
}

func (c ChainContext) GetHeader(hash libcommon.Hash, number uint64) *types.Header {
	return c.Chain.GetHeader(hash, number)
}
