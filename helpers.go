package puzgo

import (
	"errors"
	"fmt"
)

func distanceAcross(board board, x, y int) (int, error) {
	dist := 1
	curr := y
	if x > board.Width || x < 0 {
		return -1, errors.New("inavlid position")
	}
	for curr = y + 1; curr < board.Height; curr++ {
		if string(board.BoardState[x][curr]) == BLACK {
			break
		}
		dist++
	}
	return dist, nil
}

func distanceDown(board board, x, y int) (int, error) {
	dist := 1
	curr := x
	if y > board.Height || y < 0 {
		return -1, errors.New("inavlid position")
	}
	for curr = x + 1; curr < board.Width; curr++ {
		if string(board.BoardState[curr][y]) == BLACK {
			break
		}
		dist++
	}
	return dist, nil
}

func isAcrossClueNumber(board board, x, y int) (bool, error) {
	if y == 0 || string(board.BoardState[x][y-1]) == BLACK {
		dist, err := distanceAcross(board, x, y)
		if err != nil {
			return false, fmt.Errorf("error when getting number %w", err)
		}
		if dist >= 2 {
			return true, nil
		}
	}
	return false, nil
}

func isDownClueNumber(board board, x, y int) (bool, error) {
	if x == 0 || string(board.BoardState[x-1][y]) == BLACK {
		dist, err := distanceDown(board, x, y)
		if err != nil {
			return false, fmt.Errorf("error when getting number %w", err)
		}
		if dist > 2 {
			return true, nil
		}
	}
	return false, nil
}

func isBlackCell(board board, x, y int) (bool, error) {
	if x < 0 || y < 0 || x > board.Width || y > board.Height {
		return false, errors.New("invalid coordinate")
	}
	if board.BoardState[x][y] == BLACK {
		return true, nil
	}
	return false, nil
}
