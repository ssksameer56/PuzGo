package puzgo

import (
	"errors"
	"os"
)

type FileHandler struct {
	File *os.File
	Name string
}

func (f *FileHandler) ReadData(fp *os.File, size, offset int) ([]byte, error) {
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

func (f *FileHandler) VerifyChecksum() ([]byte, error) {
	return []byte{}, nil
}
