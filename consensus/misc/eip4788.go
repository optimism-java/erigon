package misc

import (
	"github.com/ledgerwatch/log/v3"

	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/optimism-java/erigon/consensus"
	"github.com/optimism-java/erigon/params"
)

func ApplyBeaconRootEip4788(parentBeaconBlockRoot *libcommon.Hash, syscall consensus.SystemCall) {
	_, err := syscall(params.BeaconRootsAddress, parentBeaconBlockRoot.Bytes())
	if err != nil {
		log.Warn("Failed to call beacon roots contract", "err", err)
	}
}
