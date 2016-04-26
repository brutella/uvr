package uvr

import (
	"github.com/brutella/canopen"
)

type Inlet struct {
	Description canopen.ObjectIndex
	Value       canopen.ObjectIndex
	State       canopen.ObjectIndex
}

func NewInlet(subIndex uint8) Inlet {
	return Inlet{
		Description: canopen.NewObjectIndex(0x2084, subIndex),
		Value:       canopen.NewObjectIndex(0x208d, subIndex),
		State:       canopen.NewObjectIndex(0x208e, subIndex),
	}
}
