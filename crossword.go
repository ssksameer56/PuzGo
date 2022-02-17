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
	flag, err = cw.parseStrings(fp)
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
	cw.Header.LowMaskChecksum, err = utils.ConvertToInt(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing lowMaskChecksum %w", err)
	}
	data, err = getData(fp, headerFormat, "highMaskChecksum")
	if err != nil {
		return false, fmt.Errorf("error while parsing highMaskChecksum %w", err)
	}
	cw.Header.HighMaskChecksum, err = utils.ConvertToInt(data) // Make sure you know if the data is LittleEndian or BigEndian
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
	cw.Header.Width, err = utils.ConvertByteToInt(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing width %w", err)
	}
	data, err = getData(fp, headerFormat, "height")
	if err != nil {
		return false, fmt.Errorf("error while parsing height %w", err)
	}
	cw.Header.Height, err = utils.ConvertByteToInt(data) // Make sure you know if the data is LittleEndian or BigEndian
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
	cw.Header.IsScrambled, err = utils.ConvertToBool(data) // Make sure you know if the data is LittleEndian or BigEndian
	if err != nil {
		return false, fmt.Errorf("error while parsing scrambledTag %w", err)
	}
	return true, nil
}

//Parses the state section from puz file
func (cw *Crossword) parseState(fp *os.File) (bool, error) {
	length := cw.Header.Height * cw.Header.Width
	cw.Board.Height = cw.Header.Height
	cw.Board.Width = cw.Header.Width
	cw.Board.BoardState = make([][]string, cw.Board.Width)
	cw.Board.Answers = make([][]string, cw.Board.Width)
	for i := 0; i < cw.Header.Width; i++ {
		cw.Board.BoardState[i] = make([]string, cw.Header.Height)
		cw.Board.Answers[i] = make([]string, cw.Header.Height)
	}
	byteArray, err := getDataWithSpecifiedLength(fp, stateFormat, "answerString", length)
	if err != nil {
		return false, fmt.Errorf("error when parsing answers section %w", err)
	}
	//Parse the correct answer state
	stringData := string(byteArray)
	num := 0
	for i := 0; i < cw.Header.Width; i++ {
		for j := 0; j < cw.Header.Height; j++ {
			cw.Board.Answers[i][j] = string(stringData[num])
			num++
		}
	}
	offset := stateFormat["answerString"].Offset + (cw.Header.Height * cw.Header.Width)
	byteArray, err = getDataWithSpecifiedOffsetLength(fp, stateFormat, "playerStateString", offset, length)
	if err != nil {
		return false, fmt.Errorf("error when parsing state section %w", err)
	}
	stringData = string(byteArray)
	num = 0
	for i := 0; i < cw.Header.Width; i++ {
		for j := 0; j < cw.Header.Height; j++ {
			cw.Board.BoardState[i][j] = string(stringData[num])
			num++
		}
	}
	//
	return true, nil
}

//Parses the clues and other strings in the puz file
func (cw *Crossword) parseStrings(fp *os.File) (bool, error) {
	offset := 0x34 + 2*(cw.Header.Height*cw.Header.Width)
	fi, _ := fp.Stat()
	length := fi.Size() - int64(offset)
	byteArray, err := getDataWithSpecifiedOffsetLength(fp, stringsFormat, "strings", offset, int(length))
	if err != nil {
		return false, fmt.Errorf("error when parsing string section %w", err)
	}
	allString, err := utils.SplitByteString(byteArray, byte(0))
	if err != nil {
		return false, fmt.Errorf("error when spliting strings %w", err)
	}
	count := 0
	cw.Header.Title = allString[count]
	count++
	cw.Header.Author = allString[count]
	count++
	cw.Header.Copyright = allString[count]
	count++

	clueNumber := 1
	var clueSlice []clueInfo
	for i := 0; i < cw.Board.Width; i++ {
		for j := 0; j < cw.Board.Height; j++ {
			blackFlag, _ := isBlackCell(cw.Board, i, j)
			if blackFlag {
				continue
			}
			acrossFlag, err := isAcrossClueNumber(cw.Board, i, j)
			if err != nil {
				cw.Clues = clueSlice
				return false, fmt.Errorf("error when assigning clue numbers %w", err)
			}
			if acrossFlag {
				clueSlice = append(clueSlice, clueInfo{
					Position: position{
						X: i,
						Y: j,
					},
					Clue:      allString[count],
					Number:    clueNumber,
					Direction: ACROSS,
				})
				count++
			}
			downFlag, _ := isDownClueNumber(cw.Board, i, j)
			if err != nil {
				cw.Clues = clueSlice
				return false, fmt.Errorf("error when assiging clue numbers %w", err)
			}
			if downFlag {
				clueSlice = append(clueSlice, clueInfo{
					Position: position{
						X: i,
						Y: j,
					},
					Clue:      allString[count],
					Number:    clueNumber,
					Direction: DOWN,
				})
				count++
			}
			if acrossFlag || downFlag {
				clueNumber++
			}
		}
	}
	cw.Header.Notes = allString[count]
	count++
	cw.Clues = clueSlice
	return true, nil
}

//Function wrapper to read data from a file. with specified field from the header
func getData(fp *os.File, format binaryFormat, binaryField string) ([]byte, error) {
	params := format[binaryField]
	data, err := utils.ReadData(fp, params.Length, params.Offset)
	if err != nil {
		return []byte{}, fmt.Errorf("error when reading %w", err)
	}
	return data, nil
}

//Get Data but when length is computed and not defined in the format
func getDataWithSpecifiedLength(fp *os.File, format binaryFormat, binaryField string, len int) ([]byte, error) {
	params := format[binaryField]
	if params.Offset != -1 {
		data, err := utils.ReadData(fp, len, params.Offset)
		if err != nil {
			return []byte{}, fmt.Errorf("error when reading %w", err)
		}
		return data, nil
	}
	return []byte{}, fmt.Errorf("invalid offset for %s", binaryField)
}

//Get data but when offset is computed and not defined in format
func getDataWithSpecifiedOffsetLength(fp *os.File, format binaryFormat, binaryField string, off int, len int) ([]byte, error) {
	params := format[binaryField]
	if params.Length == -1 {
		data, err := utils.ReadData(fp, len, off)
		if err != nil {
			return []byte{}, fmt.Errorf("error when reading %w", err)
		}
		return data, nil
	}
	return []byte{}, fmt.Errorf("invalid length for %s", binaryField)
}
