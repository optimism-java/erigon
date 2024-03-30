package raw

import (
	_ "embed"

	"github.com/optimism-java/erigon/cl/clparams"
	"github.com/optimism-java/erigon/cl/utils"
)

//go:embed testdata/state.ssz_snappy
var denebState []byte

func GetTestState() *BeaconState {
	state := New(&clparams.MainnetBeaconConfig)
	utils.DecodeSSZSnappy(state, denebState, int(clparams.DenebVersion))
	return state

}
