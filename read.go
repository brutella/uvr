package uvr

import (
	"fmt"
	"github.com/brutella/can"
	"github.com/brutella/canopen"
	"github.com/brutella/canopen/sdo"
	"time"
)

func ReadFromIndex(idx canopen.ObjectIndex, nodeID uint8, bus *can.Bus) (interface{}, error) {
	upload := sdo.Upload{
		ObjectIndex:   idx,
		RequestCobID:  uint16(SSDOClientToServer2) + uint16(nodeID),
		ResponseCobID: uint16(SSDOServerToClient2) + uint16(nodeID),
	}

	b, err := upload.Do(bus)

	if err != nil {
		return nil, err
	}

	if len(b) != 7 {
		return nil, fmt.Errorf("Invalid number of received bytes")
	}

	// isParam := b[6]&0x80
	dt := b[6] & 0x7F
	dataType := dt & 0x70
	switch dataType {
	case 0x10: // String reference
		index := parseStringIndex(b)
		return ReadStringAtIndex(index, nodeID, bus)

	case 0x20: // Bit field
		return fmt.Sprintf("Bits %b", b[0]), nil

	case 0x30:
		return parseCharacter(b)

	case 0x40: // 16-bit integer converted to a 32-bit float
		return parseInt16(b), nil
	case 0x50: // 32-bit integer converted to a 32-bit float
		return parseInt32(b), nil

		// index := parseUnitIndex(b)
		// str, err := ReadStringAtIndex(index, nodeID, bus)
		// if err != nil {
		//     return nil, err
		// }
		// return fmt.Sprintf("value %f %s", value, str), nil

	default:
		break
	}

	return nil, fmt.Errorf("Unknown data type %X", dataType)
}

func ReadStringAtIndex(idx canopen.ObjectIndex, nodeID uint8, bus *can.Bus) (string, error) {
	upload := sdo.Upload{
		ObjectIndex:   idx,
		RequestCobID:  uint16(SSDOClientToServer2) + uint16(nodeID),
		ResponseCobID: uint16(SSDOServerToClient2) + uint16(nodeID),
	}
	b, err := upload.Do(bus)
	if err != nil {
		return "", err
	}

	// b may contain invalid utf8 characters
	return printableASCIIString(b), nil
}

func printableASCIIString(b []byte) string {
	var ascii []byte
	for _, b := range b {
		if b >= 32 && b <= 126 {
			ascii = append(ascii, b)
		}
	}

	return string(ascii)
}

func parseStringIndex(b []byte) canopen.ObjectIndex {
	index := canopen.ObjectIndex{}
	index.Index.B0 = b[4]
	index.Index.B1 = b[5]
	index.SubIndex = b[0]

	return index
}

func parseCharacter(b []byte) (interface{}, error) {
	dt := b[6] & 0x7F
	value := byte(b[0])
	floatValue := float32(value)
	decimal := int(b[4] & 0x1)
	if decimal > 0 {
		floatValue = floatValue / float32(decimal*10)
	}

	switch dt {
	case 0x32:
		if value == 0 {
			return OutletStateOff, nil
		}
		return floatValue, nil
	case 0x33:
		return value, nil
	case 0x34:
		if value <= 25 {
			return string(0x41 + value), nil
		} else {
			return nil, fmt.Errorf("Invalid value %d for data type %d", value, dt)
		}
	case 0x35:
		return value * 5, nil
	case 0x36:
		if value <= 90 {
			return time.Duration(value) * time.Second, nil
		} else if value <= 107 {
			return time.Duration(value-87) * 30 * time.Second, nil
		} else if value <= 157 {
			return time.Duration(value-97) * 60 * time.Second, nil
		} else {
			return time.Duration(value-155) * 1800 * time.Second, nil
		}
	default:
		return nil, fmt.Errorf("Unsupported character data type %X", dt)
	}
}

func parseUnitIndex(b []byte) canopen.ObjectIndex {
	return canopen.NewObjectIndex(0x5002, uint8(b[5]))
}

// parseInt32 parses a 32-bit integer and converts it to a 32-bit float
func parseInt32(b []byte) float32 {
	// Bytes
	// 0: first 8 bits of int32
	// 1: second 8 bits of int32
	// 2: third 8 bits of int32
	// 3: fourth 8 bits of int32
	// 4: LSD = number of decimal places; MSB = ignored for now
	// 5: subindex of unit object at index 0x5002
	// 6: ?
	decimal := byte(b[4] & 0x1)
	value := (int32(b[3]) << 24) + (int32(b[2]) << 16) + (int32(b[1]) << 8) + int32(b[0])
	floatValue := float32(value)
	if decimal > 0 {
		floatValue = floatValue / (float32(decimal) * 10)
	}

	return floatValue
}

// parseInt16 parses a 16-bit integer and converts it to a 32-bit float
func parseInt16(b []byte) float32 {
	// Bytes
	// 0: first 8 bits of int16
	// 1: second 8 bits of int16
	// 2: ?
	// 3: ?
	// 4: LSD = number of decimal places; MSB = ignored for now
	// 5: subindex of unit object at index 0x5002
	// 6: ?
	decimal := byte(b[4] & 0x1)
	value := (int16(b[1]) << 8) + int16(b[0])
	floatValue := float32(value)
	if decimal > 0 {
		floatValue = floatValue / (float32(decimal) * 10)
	}

	return floatValue
}
