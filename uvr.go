package uvr

const (
	TPDODigitalVariables     uint16 = 0x180
	TPDOAnalogVariables01_04 uint16 = 0x200
	TPDOAnalogVariables17_20 uint16 = 0x240
	TPDOAnalogVariables05_08 uint16 = 0x280
	TPDOAnalogVariables21_24 uint16 = 0x2C0
	TPDOAnalogVariables09_12 uint16 = 0x300
	TPDOAnalogVariables25_28 uint16 = 0x340
	TPDOAnalogVariables13_16 uint16 = 0x380
	TPDOAnalogVariables29_32 uint16 = 0x3C0

	MPDOClientServerConnManagement uint16 = 0x400
	MPDOLogging                    uint16 = 0x480
	MPDOUnits                      uint16 = 0x480

	SSDOServerToClient1 uint16 = 0x580
	SSDOServerToClient2 uint16 = 0x5C0
	SSDOClientToServer1 uint16 = 0x600
	SSDOClientToServer2 uint16 = 0x640
)

const (
	OutletModeAuto    = "AUTO"
	OutletModeTimed   = "SCHZ"
	OutletModeManuell = "HAND"

	OutletStateOn  = "EIN"
	OutletStateOff = "AUS"

	DescriptionUnused = "-----"
	OutletModeUnused  = ""
)
