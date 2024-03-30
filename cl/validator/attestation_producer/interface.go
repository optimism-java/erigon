package attestation_producer

import (
	"github.com/optimism-java/erigon/cl/cltypes/solid"
	"github.com/optimism-java/erigon/cl/phase1/core/state"
)

type AttestationDataProducer interface {
	ProduceAndCacheAttestationData(baseState *state.CachingBeaconState, slot uint64, committeeIndex uint64) (solid.AttestationData, error)
}
