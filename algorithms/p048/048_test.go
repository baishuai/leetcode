package p048

import (
	"testing"
)

func test(t *testing.T, matrix [][]int, exp [][]int) {
	rotate(matrix)
	if len(matrix) != len(exp) {
		t.Fatal("error answer length")
	}
	for i, v := range matrix {
		expv := exp[i]
		if len(v) != len(expv) {
			t.Fatal("error answer length")
		}
		for j := 0; j < len(v); j++ {
			if v[j] != expv[j] {
				t.Fatal("error answer value", "i", i, "j", j, v[j], expv[j])
			}
		}
	}
}

func TestExample(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	exp := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	test(t, matrix, exp)
}
