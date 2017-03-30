package p130

import "testing"

func TestExample(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
}

func TestExample1(t *testing.T) {
	board := [][]byte{
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
	}
	solve(board)
}
