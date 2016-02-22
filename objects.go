package uvr

import (
	"github.com/brutella/canopen"
)

type Input struct {
	Description canopen.ObjectIndex
	Value       canopen.ObjectIndex
	State       canopen.ObjectIndex
}

func NewTempCollector() Input {
	return Input{
		Description: canopen.NewObjectIndex(0x2084, 0x1),
		Value:       canopen.NewObjectIndex(0x208d, 0x1),
		State:       canopen.NewObjectIndex(0x208e, 0x1),
	}
}

func NewTempHeater1() Input {
	return Input{
		Description: canopen.NewObjectIndex(0x2084, 0x2),
		Value:       canopen.NewObjectIndex(0x208d, 0x2),
		State:       canopen.NewObjectIndex(0x208e, 0x2),
	}
}

func NewTempHeater3() Input {
	return Input{
		Description: canopen.NewObjectIndex(0x2084, 0x3),
		Value:       canopen.NewObjectIndex(0x208d, 0x3),
		State:       canopen.NewObjectIndex(0x208e, 0x3),
	}
}

type Output struct {
	Description canopen.ObjectIndex
	StartDelay  canopen.ObjectIndex
	RunOnTime   canopen.ObjectIndex
	Mode        canopen.ObjectIndex
	State       canopen.ObjectIndex
	SpeedStage  canopen.ObjectIndex
}

func NewSolarPump1() Output {
	return Output{
		Description: canopen.NewObjectIndex(0x20a5, 0x1),
		StartDelay:  canopen.NewObjectIndex(0x20a3, 0x1),
		RunOnTime:   canopen.NewObjectIndex(0x20a4, 0x1),
		Mode:        canopen.NewObjectIndex(0x20a1, 0x1),
		State:       canopen.NewObjectIndex(0x20aa, 0x1),
		SpeedStage:  canopen.NewObjectIndex(0x20ab, 0x1),
	}
}
func NewHeating1() Output {
	return Output{
		Description: canopen.NewObjectIndex(0x20a5, 0x3),
		StartDelay:  canopen.NewObjectIndex(0x20a3, 0x3),
		RunOnTime:   canopen.NewObjectIndex(0x20a4, 0x3),
		Mode:        canopen.NewObjectIndex(0x20a1, 0x3),
		State:       canopen.NewObjectIndex(0x20aa, 0x3),
		SpeedStage:  canopen.NewObjectIndex(0x20ab, 0x3),
	}
}

func NewMixerHeating1() Output {
	return Output{
		Description: canopen.NewObjectIndex(0x20a5, 0x8),
		StartDelay:  canopen.NewObjectIndex(0x20a3, 0x8),
		RunOnTime:   canopen.NewObjectIndex(0x20a4, 0x8),
		Mode:        canopen.NewObjectIndex(0x20a1, 0x8),
		State:       canopen.NewObjectIndex(0x20aa, 0x8),
		SpeedStage:  canopen.NewObjectIndex(0x20ab, 0x8),
	}
}

func NewMixerHeating2() Output {
	return Output{
		Description: canopen.NewObjectIndex(0x20a5, 0xa),
		StartDelay:  canopen.NewObjectIndex(0x20a3, 0xa),
		RunOnTime:   canopen.NewObjectIndex(0x20a4, 0xa),
		Mode:        canopen.NewObjectIndex(0x20a1, 0xa),
		State:       canopen.NewObjectIndex(0x20aa, 0xa),
		SpeedStage:  canopen.NewObjectIndex(0x20ab, 0xa),
	}
}
