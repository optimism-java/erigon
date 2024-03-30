package handler

import (
	"fmt"
	"net/http"

	"github.com/ledgerwatch/erigon-lib/common"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/optimism-java/erigon/cl/beacon/beaconhttp"
	"github.com/optimism-java/erigon/cl/utils"
)

type genesisResponse struct {
	GenesisTime          uint64           `json:"genesis_time,string"`
	GenesisValidatorRoot common.Hash      `json:"genesis_validators_root"`
	GenesisForkVersion   libcommon.Bytes4 `json:"genesis_fork_version"`
}

func (a *ApiHandler) GetEthV1BeaconGenesis(w http.ResponseWriter, r *http.Request) (*beaconhttp.BeaconResponse, error) {
	if a.genesisCfg == nil {
		return nil, beaconhttp.NewEndpointError(http.StatusNotFound, fmt.Errorf("Genesis Config is missing"))
	}

	return newBeaconResponse(&genesisResponse{
		GenesisTime:          a.genesisCfg.GenesisTime,
		GenesisValidatorRoot: a.genesisCfg.GenesisValidatorRoot,
		GenesisForkVersion:   utils.Uint32ToBytes4(uint32(a.beaconChainCfg.GenesisForkVersion)),
	}), nil
}
