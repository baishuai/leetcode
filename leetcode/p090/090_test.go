package p090

import "testing"

func test(t *testing.T, nums []int, exp [][]int) {
	ans := subsetsWithDup(nums)
	if len(ans) != len(exp) {
		t.Fatal("error answer length", len(ans))
	}
}

func TestExample(t *testing.T) {
	test(t, []int{1, 2, 2}, [][]int{
		{2},
		{1},
		{1, 2, 2},
		{2, 2},
		{1, 2},
		{},
	})
}

func TestExtra0(t *testing.T) {
	test(t, []int{1, 2, 2, 3, 3}, [][]int{
		{1, 2, 2, 3, 3},
		{1, 2, 2, 3},
		{1, 2, 2},
		{1, 2, 3, 3},
		{1, 2, 3},
		{1, 2},
		{1, 3, 3},
		{1, 3},
		{1},
		{2, 2, 3, 3},
		{2, 2, 3},
		{2, 2},
		{2, 3, 3},
		{2, 3},
		{2},
		{3, 3},
		{3},
		{},
	})
}
