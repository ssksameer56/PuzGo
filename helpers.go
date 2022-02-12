package puzgo

import (
	"errors"
)

func distanceAcross(board Board, x, y int) (int, error) {
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

func distanceDown(board Board, x, y int) (int, error) {
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

func isBlack(board Board, x, y int) (bool, error) {
	if x > board.Width || y > board.Height || x < 0 || y < 0 {
		return false, errors.New("inavlid position")
	}
	if string(board.BoardState[x][y]) == BLACK {
		return false, nil
	}
	return true, nil
}

func isBlank(board Board, x, y int) (bool, error) {
	if x > board.Width || y > board.Height || x < 0 || y < 0 {
		return false, errors.New("inavlid position")
	}
	if string(board.BoardState[x][y]) == BLANK {
		return false, nil
	}
	return true, nil
}
