package p079

import "testing"

func TestExample(t *testing.T) {
	ans := exist([][]byte{{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}, "ABCCED")
	if !ans {
		t.Fatal("error")
	}
}
