package handler

import (
	"github.com/optimism-java/erigon/cl/beacon/beaconhttp"
)

func newBeaconResponse(data any) *beaconhttp.BeaconResponse {
	return beaconhttp.NewBeaconResponse(data)
}
