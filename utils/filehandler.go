package utils

import (
	"errors"
	"os"
)

//Read Data from a file. Custom function created to handle length and offset directly.
//Returns the data and error if any.
//Returns error if the data read is less than expected `size` value
func ReadData(fp *os.File, size, offset int) ([]byte, error) {
	data := make([]byte, size)
	len, err := fp.ReadAt(data, int64(offset))
	if err != nil {
		return data, err
	}
	if len < size {
		return []byte{}, errors.New("could not read complete data")
	}
	return data, err
}

//Function to verify the checksum passed with the one in file
func VerifyChecksum() ([]byte, error) {
	return []byte{}, nil
}
