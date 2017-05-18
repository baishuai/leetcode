package p036

import "testing"

func test(t *testing.T, board [][]byte, exp bool) {
	if isValidSudoku(board) != exp {
		t.Fatal("error answer")
	}
}

func TestExample(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	test(t, board, true)
}

func TestBitset(t *testing.T) {
	b := bit9Set(0)
	if b.containsOrInsert(5) {
		t.Error("error contains")
	}
	if !b.containsOrInsert(5) {
		t.Error("error contains")
	}
}
