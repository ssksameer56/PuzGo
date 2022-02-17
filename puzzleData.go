package puzgo

//Constants used in crossword

var BLACK string = "." //Denotes empty cell according to puz
var BLANK string = "-" //Denotes a black cell according to puz
var ACROSS string = "across"
var DOWN string = "down"

//Struct to map position in a crossword
type position struct {
	X int
	Y int
}

//Struct to store a clue(number, position and string)
type clueInfo struct {
	Position  position
	Number    int
	Clue      string
	Direction string
}

//Struct to keep the board state(empty,filled,black) cells
type board struct {
	Height     int
	Width      int
	BoardState [][]string
	Answers    [][]string
}

//Store information about crossword from header and strings section in puz file
type crosswordInfo struct {
	Title             string `binary:"strings"`
	Author            string `binary:"strings"`
	Copyright         string `binary:"strings"`
	Notes             string `binary:"strings"`
	Version           string `binary:"version"`
	Width             int    `binary:"width"`
	Height            int    `binary:"height"`
	NoOfClues         int16  `binary:"noOfClues"`
	IsScrambled       bool   `binary:"scrambledTag"`
	CibChecksum       int16  `binary:"strings"`
	LowMaskChecksum   int    `binary:"lowMaskChecksum"`
	HighMaskChecksum  int    `binary:"highMaskChecksum"`
	ScrambledChecksum int16  `binary:"scrambledChecksum"`
	Checksum          int16  `binary:"checksum"`
}

//Exported struct for the crossword
type Crossword struct {
	Header   crosswordInfo
	Board    board
	Clues    []clueInfo
	filePath string
}
