package uvr

import (
	"github.com/brutella/can"
	"github.com/brutella/canopen"
	"github.com/brutella/canopen/sdo"
)

func WriteToIndex(idx canopen.ObjectIndex, b []byte, nodeID uint8, bus *can.Bus) error {
	download := sdo.Download{
		ObjectIndex:   idx,
		RequestCobID:  uint16(SSDOClientToServer2) + uint16(nodeID),
		ResponseCobID: uint16(SSDOServerToClient2) + uint16(nodeID),
		Data:          b,
	}

	return download.Do(bus)
}
