package p120

import "testing"

func TestExample(t *testing.T) {
	ans := minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	})
	if ans != 11 {
		t.Fail()
	}
}
