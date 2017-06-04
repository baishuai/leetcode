package p054

import "testing"

func test(t *testing.T, mat [][]int, exp []int) {
	ans := spiralOrder(mat)
	if len(ans) != len(exp) {
		t.Fatal("error answer length")
	}
	for i, v := range exp {
		if v != ans[i] {
			t.Error("error answer value", i, v, exp[i])
		}
	}
}

func TestExample(t *testing.T) {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	test(t, mat, []int{1, 2, 3, 6, 9, 8, 7, 4, 5})
}

func TestExtra0(t *testing.T) {
	mat := [][]int{
		{},
	}
	test(t, mat, []int{})
}

func TestExtra1(t *testing.T) {
	mat := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	test(t, mat, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})
}
