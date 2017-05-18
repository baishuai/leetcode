package p154

import "testing"

func TestExample(t *testing.T) {
	ans := findMin([]int{4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 7, 0, 1, 2, 4, 4, 4, 4, 4, 4, 4, 4})
	if ans != 0 {
		t.Fail()
	}
}
