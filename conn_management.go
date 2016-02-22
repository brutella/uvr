package uvr

import (
	"fmt"
	"github.com/brutella/can"
	"github.com/brutella/canopen"
	"time"
)

func Disconnect(serverID canopen.NodeID, clientID canopen.NodeID, bus *can.Bus) error {
    b := []byte{
			0x80 + byte(serverID),
			0x01, 0x1F,
			0x00,
			byte(serverID),
			byte(clientID),
			0x80,
			0x12,
	}
        
    return sendConnManagementData(b, serverID, clientID, bus)
}

func Connect(serverID canopen.NodeID, clientID canopen.NodeID, bus *can.Bus) error {
    b := []byte{
			0x80 + byte(serverID),
			0x00, 0x1F,
			0x00,
			byte(serverID),
			byte(clientID),
			0x80,
			0x12,
	}
        
    return sendConnManagementData(b, serverID, clientID, bus)
}

func sendConnManagementData(b []byte, serverID canopen.NodeID, clientID canopen.NodeID, bus *can.Bus) error {
	c := &canopen.Client{bus, time.Second * 2}
	frm := canopen.Frame{
		CobID: uint16(MPDOClientServerConnManagement) + uint16(clientID),
		Data: b,
	}
    
	respCobID := uint32(MPDOClientServerConnManagement) + uint32(serverID)
	req := canopen.NewRequest(frm, respCobID)
	resp, err := c.Do(req)

	if err != nil {
		return err
	}

	frm = resp.Frame

	if b0 := frm.Data[0]; b0 != 0x80+byte(clientID) {
		return fmt.Errorf("Invalid MPDO address %v\n", b0)
	}

	if b4, b5 := frm.Data[4], frm.Data[5]; b4 != 0x40+byte(clientID) || b5 != 0x06 {
		return fmt.Errorf("Invalid 0x640 + client id %X %X\n", b5, b4)
	}

	if b7 := frm.Data[7]; b7 != 0x00 {
		return fmt.Errorf("Invalid byte 7 %X", b7)
	}

	return nil
}