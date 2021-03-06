package p198

import "testing"

func test(t *testing.T, nums []int, exp int) {
	ans := rob(nums)
	if ans != exp {
		t.Fatal(ans, "exp", exp)
	}
}

func TestExample0(t *testing.T) {
	test(t, []int{}, 0)
	test(t, nil, 0)
}

func TestExample1(t *testing.T) {
	test(t, []int{1}, 1)
	test(t, []int{1, 2}, 2)
	test(t, []int{1, 2, 0}, 2)
}
