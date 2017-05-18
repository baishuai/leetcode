package p136

import "testing"

func test(t *testing.T, nums []int, exp int) {
	if ans := singleNumber(nums); ans != exp {
		t.Fatal("error", ans, exp)
	}

	if ans := singleNumber_slow(nums); ans != exp {
		t.Fatal("error", ans, exp)
	}
}

func TestEXample0(t *testing.T) {
	test(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2, 1}, 8)
}
