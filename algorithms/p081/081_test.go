package p081

import "testing"

func test(t *testing.T, nums []int, target int, exp bool) {
	if search(nums, target) != exp {
		t.Fatal("error test")
	}
}

func TestExtra0(t *testing.T) {
	test(t, []int{1, 1, 3}, 2, false)
}

func TestExtra1(t *testing.T) {
	test(t, []int{1, 3, 1, 1, 1}, 3, true)
}
