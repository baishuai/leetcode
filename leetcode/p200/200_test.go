package p200

import "testing"

func TestExample(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	if numIslands(grid) != 3 {
		t.FailNow()
	}
}
