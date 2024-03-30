package stagedsynctest

import (
	"github.com/ledgerwatch/erigon-lib/chain"
	"github.com/optimism-java/erigon/params"
	"github.com/optimism-java/erigon/polygon/bor/borcfg"
)

func BorDevnetChainConfigWithNoBlockSealDelays() *chain.Config {
	// take care not to mutate global var (shallow copy)
	chainConfigCopy := *params.BorDevnetChainConfig
	borConfigCopy := *chainConfigCopy.Bor.(*borcfg.BorConfig)
	borConfigCopy.Period = map[string]uint64{
		"0": 0,
	}
	borConfigCopy.ProducerDelay = map[string]uint64{
		"0": 0,
	}
	chainConfigCopy.Bor = &borConfigCopy
	return &chainConfigCopy
}
