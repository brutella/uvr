package uvr

import (
	"fmt"
	"github.com/brutella/can"
	"github.com/brutella/canopen"
	"time"
)

func ConnectTo(toNodeID canopen.NodeID, nodeID canopen.NodeID, bus *can.Bus) error {
	c := &canopen.Client{bus, time.Second * 2}
	frm := canopen.Frame{
		CobID: uint16(MPDOClientServerAddressing) + uint16(nodeID),
		Data: []byte{
			0x80 + byte(toNodeID),
			0x00, 0x1F,
			0x00,
			byte(toNodeID),
			byte(nodeID),
			0x80,
			0x12,
		},
	}

	respCobID := uint32(MPDOClientServerAddressing) + uint32(toNodeID)
	req := canopen.NewRequest(frm, respCobID)
	resp, err := c.Do(req)

	if err != nil {
		return err
	}

	frm = resp.Frame

	if b0 := frm.Data[0]; b0 != 0x80+byte(nodeID) {
		return fmt.Errorf("Invalid MPDO address %v\n", b0)
	}

	if b4, b5 := frm.Data[4], frm.Data[5]; b4 != 0x40+byte(nodeID) || b5 != 0x06 {
		return fmt.Errorf("Invalid 0x640 + client id %X %X\n", b5, b4)
	}

	if b7 := frm.Data[7]; b7 != 0x00 {
		return fmt.Errorf("Invalid byte 7 %X", b7)
	}

	return nil
}
