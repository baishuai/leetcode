package p085

import "testing"

func TestExample(t *testing.T) {
	mat := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	if res := maximalRectangle(mat); res != 6 {
		t.Fail()
	}
}
