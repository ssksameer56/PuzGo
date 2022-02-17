package utils

import (
	"encoding/binary"
	"errors"
	"strings"
)

//Convert slice of bytes to int16 or short
func ConvertToInt16(data []byte) (int16, error) {
	if len(data) != 2 {
		return -1, errors.New("slice must be 2 byte long for int16")
	}
	intData := binary.LittleEndian.Uint16(data)
	return int16(intData), nil
}

func ConvertToInt(data []byte) (int, error) {
	if len(data) != 4 {
		return -1, errors.New("slice must be 2 byte long for int16")
	}
	intData := binary.LittleEndian.Uint32(data)
	return int(intData), nil
}

func ConvertByteToInt(data []byte) (int, error) {
	if len(data) != 1 {
		return -1, errors.New("slice must be 2 byte long for int16")
	}
	return int(data[0]), nil
}

func ConvertToString(data []byte) (string, error) {
	var newData []byte
	for _, x := range data {
		if x != 0x00 {
			newData = append(newData, x)
		}

	}
	return string(newData), nil
}

func SplitByteString(data []byte, delim byte) ([]string, error) {
	var stringsData []string
	masterString := string(data)
	stringsData = strings.Split(masterString, string(delim))
	var stringsFormatted []string
	for _, x := range stringsData {
		xFormatted := strings.Trim(x, " ")
		stringsFormatted = append(stringsFormatted, string(xFormatted))
	}
	return stringsFormatted, nil
}

func ConvertToBool(data []byte) (bool, error) {
	if len(data) == 2 {
		intData, _ := ConvertToInt16(data)
		if intData == 0 {
			return false, nil
		}
		return true, nil
	} else if len(data) == 4 {
		intData, _ := ConvertToInt(data)
		if intData == 0 {
			return false, nil
		}
		return true, nil
	} else if len(data) == 1 {
		intData, _ := ConvertByteToInt(data)
		if intData == 0 {
			return false, nil
		}
		return true, nil
	}
	return false, errors.New("invalid length of byte slice. must be 1,2 or 4")
}
