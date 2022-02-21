package uvr

import (
	"github.com/brutella/canopen"
)

type AnalogOutlet struct {
	Description canopen.ObjectIndex
	Mode        canopen.ObjectIndex
	Value       canopen.ObjectIndex
}

func NewAnalogOutlet(subIndex uint8) AnalogOutlet {
	return AnalogOutlet{
		Description: canopen.NewObjectIndex(0x20c1, subIndex),
		Mode:        canopen.NewObjectIndex(0x20c5, subIndex),
		Value:       canopen.NewObjectIndex(0x20c9, subIndex),
	}
}
