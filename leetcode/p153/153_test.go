package p153

import "testing"

func TestExample(t *testing.T) {
	ans := findMin([]int{4, 5, 6, 7, 0, 1, 2})
	if ans != 0 {
		t.Fail()
	}
}
