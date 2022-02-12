package puzgo

import (
	"errors"
	"fmt"
	"os"

	"github.com/ssksameer56/PuzGo/utils"
)

//Returns an array with width and height of the crossword
func (cw *Crossword) GetSize() ([]int, error) {
	return []int{cw.Header.Height, cw.Header.Width}, nil
}

//Parses the puz file to generate the sections of crossword - clues, state and info
//Returns true if successful
//Returns false if any error occured
func (cw *Crossword) ParseCrossword() (bool, error) {
	fp, err := os.Open(cw.filePath)
	if err != nil {
		return false, fmt.Errorf("error while opening file: %w", err)
	}
	defer fp.Close()

	flag, err := cw.parseHeader(fp)
	if err != nil || !flag {
		return false, fmt.Errorf("error while parsing header: %w", err)
	}
	flag, err = cw.parseState(fp)
	if err != nil || !flag {
		return false, fmt.Errorf("error while parsing state: %w", err)
	}
	flag, err = cw.parseClues(fp)
	if err != nil || !flag {
		return false, fmt.Errorf("error while parsing clues: %w", err)
	}
	return true, nil
}

//Initializer to create a new crossword struct with the path of puz file
func NewCrossword(path string) (*Crossword, error) {
	if path == "" {
		return &Crossword{}, errors.New("empty file provided")
	}
	return &Crossword{
		filePath: path,
	}, nil
}

//Parses the header from puz file
func (cw *Crossword) parseHeader(fp *os.File) (bool, error) {
	cw.Header = crosswordInfo{}
	var data []byte
	var err error
	data, err = getData(fp, headerFormat, "checksum")
	if err != nil {
		return false, fmt.Errorf("error while parsing checksum %w", err)
	}
	cw.Header.Checksum, err = utils.ConvertToInt16(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing checksum %w", err)
	}
	data, err = getData(fp, headerFormat, "cibChecksum")
	if err != nil {
		return false, fmt.Errorf("error while parsing cibChecksum %w", err)
	}
	cw.Header.Checksum, err = utils.ConvertToInt16(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing cibChecksum %w", err)
	}
	data, err = getData(fp, headerFormat, "lowMaskChecksum")
	if err != nil {
		return false, fmt.Errorf("error while parsing lowMaskChecksum %w", err)
	}
	cw.Header.LowMaskChecksum, err = utils.ConvertToInt16(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing lowMaskChecksum %w", err)
	}
	data, err = getData(fp, headerFormat, "highMaskChecksum")
	if err != nil {
		return false, fmt.Errorf("error while parsing highMaskChecksum %w", err)
	}
	cw.Header.HighMaskChecksum, err = utils.ConvertToInt16(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing highMaskChecksum %w", err)
	}
	data, err = getData(fp, headerFormat, "version")
	if err != nil {
		return false, fmt.Errorf("error while parsing version %w", err)
	}
	cw.Header.Version, err = utils.ConvertToString(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing version %w", err)
	}
	data, err = getData(fp, headerFormat, "scrambledChecksum")
	if err != nil {
		return false, fmt.Errorf("error while parsing scrambledChecksum %w", err)
	}
	cw.Header.ScrambledChecksum, err = utils.ConvertToInt16(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing scrambledChecksum %w", err)
	}
	data, err = getData(fp, headerFormat, "width")
	if err != nil {
		return false, fmt.Errorf("error while parsing width %w", err)
	}
	cw.Header.Width, err = utils.ConvertToInt(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing width %w", err)
	}
	data, err = getData(fp, headerFormat, "height")
	if err != nil {
		return false, fmt.Errorf("error while parsing height %w", err)
	}
	cw.Header.Width, err = utils.ConvertToInt(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing height %w", err)
	}
	data, err = getData(fp, headerFormat, "noOfClues")
	if err != nil {
		return false, fmt.Errorf("error while parsing noOfClues %w", err)
	}
	cw.Header.NoOfClues, err = utils.ConvertToInt16(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing noOfClues %w", err)
	}
	data, err = getData(fp, headerFormat, "scrambledTag")
	if err != nil {
		return false, fmt.Errorf("error while parsing scrambledTag %w", err)
	}
	cw.Header.isScrambled, err = utils.ConvertToBool(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, err
	}
	return true, nil
}

//Parses the clues section from puz file
func (cw *Crossword) parseClues(fp *os.File) (bool, error) {
	return true, nil
}

//Parses the state in the puz file
func (cw *Crossword) parseState(fp *os.File) (bool, error) {
	return true, nil
}

func getData(fp *os.File, format binaryFormat, binaryField string) ([]byte, error) {
	params := format[binaryField]
	data, err := utils.ReadData(fp, params.Length, params.Offset)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}