package puzgo

const (
	BLANK  = "."
	BLACK  = "-"
	ACROSS = "across"
	DOWN   = "down"
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ClueInfo struct {
	Position Position `json:"position"`
	Number   int      `json:"number"`
	Clue     string   `json:"clue"`
}

type Board struct {
	Height     int      `json:"height"`
	Width      int      `json:"width"`
	BoardState [][]byte `json:"boardState"`
}

type CrosswordInfo struct {
	Title        string
	Author       string
	Copyright    string
	Notes        string
	Version      string
	Width        int
	Height       int
	NoOfClues    int
	ScrambledTag int8
}

type Crossword struct {
	Header CrosswordInfo
	Board  Board
	Clues  []ClueInfo
}
