package utils

import (
	"encoding/binary"
	"strings"
)

//Convert slice of bytes to int16 or short
func ConvertToInt16(data []byte) (int16, error) {
	intData := binary.LittleEndian.Uint32(data)
	return int16(intData), nil
}

func ConvertToInt(data []byte) (int, error) {
	intData := binary.LittleEndian.Uint32(data)
	return int(intData), nil
}

func ConvertToString(data []byte) (string, error) {
	return string(data), nil
}

func SplitByteString(data []byte, delim byte) ([]string, error) {
	var stringsData []string
	masterString := string(data)
	stringsData = strings.Split(masterString, string(delim))
	return stringsData, nil
}

func ConvertToBool(data []byte) (bool, error) {
	intData := binary.LittleEndian.Uint32(data)
	if intData == 0 {
		return false, nil
	}
	return true, nil
}
