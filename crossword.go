package puzgo

import (
	"errors"
)

type Crossword2 struct {
	crossword Crossword
	filepath  string
}

func (cw *Crossword2) GetSize() ([]int, error) {
	return []int{cw.crossword.Header.Height, cw.crossword.Board.Width}, nil
}

func (cw *Crossword2) ParseCrossword() (bool, error) {
	return true, nil
}

func NewCrossword(path string) (*Crossword2, error) {
	if path == "" {
		return nil, errors.New("empty file path")
	}
	return &Crossword2{
		filepath:  path,
		crossword: Crossword{},
	}, nil
}
