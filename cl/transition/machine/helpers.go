package machine

import (
	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/optimism-java/erigon/cl/abstract"
	"github.com/optimism-java/erigon/cl/cltypes"
	"github.com/optimism-java/erigon/cl/phase1/core/state"
)

func executionEnabled(s abstract.BeaconState, payload *cltypes.Eth1Block) bool {
	return (!state.IsMergeTransitionComplete(s) && payload.BlockHash != common.Hash{}) || state.IsMergeTransitionComplete(s)
}
