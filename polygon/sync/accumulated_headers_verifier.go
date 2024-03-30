package sync

import (
	"bytes"
	"fmt"

	"github.com/optimism-java/erigon/core/types"
	"github.com/optimism-java/erigon/polygon/bor"
	"github.com/optimism-java/erigon/polygon/heimdall"
)

type AccumulatedHeadersVerifier func(waypoint heimdall.Waypoint, headers []*types.Header) error

func VerifyAccumulatedHeaders(waypoint heimdall.Waypoint, headers []*types.Header) error {
	rootHash, err := bor.ComputeHeadersRootHash(headers)
	if err != nil {
		return fmt.Errorf("VerifyAccumulatedHeaders: failed to compute headers root hash %w", err)
	}
	if !bytes.Equal(rootHash, waypoint.RootHash().Bytes()) {
		return fmt.Errorf("VerifyAccumulatedHeaders: bad headers root hash")
	}
	return nil
}
