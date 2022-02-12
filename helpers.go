package puzgo

import (
	"errors"
	"fmt"
)

func distanceAcross(board board, x, y int) (int, error) {
	dist := 0
	curr := y
	if x > board.Width || y < 0 {
		return -1, errors.New("inavlid position")
	}
	for curr = y + 1; curr < board.Width; curr++ {
		if string(board.BoardState[x][curr]) == BLACK || string(board.BoardState[x][curr]) == BLANK {
			break
		}
		dist++
	}
	return dist, nil
}

func distanceDown(board board, x, y int) (int, error) {
	dist := 0
	curr := x
	if y > board.Height || y < 0 {
		return -1, errors.New("inavlid position")
	}
	for curr = x + 1; curr < board.Height; curr++ {
		if string(board.BoardState[curr][x]) == BLACK || string(board.BoardState[curr][x]) == BLANK {
			break
		}
		dist++
	}
	return dist, nil
}

func isAcrossClueNumber(board board, x, y int) (bool, error) {
	if x == 0 || string(board.BoardState[x][y]) == BLANK {
		dist, err := distanceAcross(board, x, y)
		if err != nil {
			return false, fmt.Errorf("error when getting number %w", err)
		}
		if dist >= 2 && dist < (board.Width-x) {
			return true, nil
		}
	}
	return false, nil
}

func isDownClueNumber(board board, x, y int) (bool, error) {
	if y == 0 || string(board.BoardState[x][y]) == BLANK {
		dist, err := distanceDown(board, x, y)
		if err != nil {
			return false, fmt.Errorf("error when getting number %w", err)
		}
		if dist > 2 && dist < (board.Height-y) {
			return true, nil
		}
	}
	return false, nil
}
