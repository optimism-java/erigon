package portal

import (
	"github.com/ledgerwatch/erigon-lib/types/clonable"
	"github.com/ledgerwatch/erigon-lib/types/ssz"
	"github.com/optimism-java/erigon/cl/clparams"
	"github.com/optimism-java/erigon/cl/cltypes"
	"github.com/optimism-java/erigon/cl/merkle_tree"
)

var (
	Bellatrix = []byte{0x0, 0x0, 0x0, 0x0}
	Capella   = []byte{0xbb, 0xa4, 0xda, 0x96}
)

func ForkVersion(digest []byte) clparams.StateVersion {
	if digest[0] == 0xbb && digest[1] == 0xa4 && digest[2] == 0xda && digest[3] == 0x96 {
		return clparams.CapellaVersion
	}
	return clparams.BellatrixVersion
}

func ForkDigest(version clparams.StateVersion) []byte {
	if version < clparams.CapellaVersion {
		return Bellatrix
	}
	return Capella
}

func DecodeDynamicListForkedObject[T ssz.Unmarshaler](bytes []byte, start, end uint32, max uint64) ([]T, error) {
	if start > end || len(bytes) < int(end) {
		return nil, ssz.ErrBadOffset
	}
	buf := bytes[start:end]
	var elementsNum, currentOffset uint32
	if len(buf) > 4 {
		currentOffset = ssz.DecodeOffset(buf)
		elementsNum = currentOffset / 4
	}
	inPos := 4
	if uint64(elementsNum) > max {
		return nil, ssz.ErrTooBigList
	}
	objs := make([]T, elementsNum)
	for i := range objs {
		endOffset := uint32(len(buf))
		if i != len(objs)-1 {
			if len(buf[inPos:]) < 4 {
				return nil, ssz.ErrLowBufferSize
			}
			endOffset = ssz.DecodeOffset(buf[inPos:])
		}
		inPos += 4
		if endOffset < currentOffset || len(buf) < int(endOffset) {
			return nil, ssz.ErrBadOffset
		}
		objs[i] = objs[i].Clone().(T)
		if err := objs[i].DecodeSSZ(buf[currentOffset:endOffset], int(ForkVersion(buf[currentOffset:currentOffset+4]))); err != nil {
			return nil, err
		}
		currentOffset = endOffset
	}
	return objs, nil
}

type ForkedLightClientUpdate struct {
	LightClientUpdate *cltypes.LightClientUpdate
	Version           clparams.StateVersion
}

func NewForkedLightClientUpdate(version clparams.StateVersion) *ForkedLightClientUpdate {
	return &ForkedLightClientUpdate{
		LightClientUpdate: cltypes.NewLightClientUpdate(version),
		Version:           version,
	}
}

func (flgu *ForkedLightClientUpdate) EncodeSSZ(buf []byte) ([]byte, error) {
	dst := buf
	dst = append(dst, ForkDigest(flgu.Version)...)
	dst, err := flgu.LightClientUpdate.EncodeSSZ(dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func (flgu *ForkedLightClientUpdate) EncodingSizeSSZ() int {
	return len(ForkDigest(flgu.Version)) + flgu.LightClientUpdate.EncodingSizeSSZ()
}

func (flgu *ForkedLightClientUpdate) DecodeSSZ(buf []byte, version int) error {
	flgu.Version = clparams.StateVersion(version)
	flgu.LightClientUpdate = cltypes.NewLightClientUpdate(flgu.Version)
	err := flgu.LightClientUpdate.DecodeSSZ(buf[4:], version)
	if err != nil {
		return err
	}
	return nil
}

func (flgu *ForkedLightClientUpdate) HashSSZ() ([32]byte, error) {
	return merkle_tree.HashTreeRoot(ForkDigest(flgu.Version), flgu.LightClientUpdate)
}

func (flgu *ForkedLightClientUpdate) Clone() clonable.Clonable {
	return &ForkedLightClientUpdate{}
}

type ForkedLightClientUpdateRange []*ForkedLightClientUpdate
