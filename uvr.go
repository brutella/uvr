package uvr

import (
	"github.com/brutella/canopen"
)

const (
	TPDODigitalVariables     canopen.FunctionCode = 0x180
	TPDOAnalogVariables01_04 canopen.FunctionCode = 0x200
	TPDOAnalogVariables17_20 canopen.FunctionCode = 0x240
	TPDOAnalogVariables05_08 canopen.FunctionCode = 0x280
	TPDOAnalogVariables21_24 canopen.FunctionCode = 0x2C0
	TPDOAnalogVariables09_12 canopen.FunctionCode = 0x300
	TPDOAnalogVariables25_28 canopen.FunctionCode = 0x340
	TPDOAnalogVariables13_16 canopen.FunctionCode = 0x380
	TPDOAnalogVariables29_32 canopen.FunctionCode = 0x3C0

	MPDOClientServerConnManagement canopen.FunctionCode = 0x400
	MPDOLogging                canopen.FunctionCode = 0x480
	MPDOUnits                  canopen.FunctionCode = 0x480

	SSDOServerToClient1 canopen.FunctionCode = 0x580
	SSDOServerToClient2 canopen.FunctionCode = 0x5C0
	SSDOClientToServer1 canopen.FunctionCode = 0x600
	SSDOClientToServer2 canopen.FunctionCode = 0x640
)
