package p041

import (
	"testing"
)

func test(t *testing.T, nums []int, exp int) {
	if ans := firstMissingPositive(nums); ans != exp {
		t.Fatal("error answer", ans)
	}
}

func TestExample1(t *testing.T) {
	test(t, []int{1, 2, 0}, 3)
}

func TestExample2(t *testing.T) {
	test(t, []int{1, 2, 3}, 4)
}

func TestExample3(t *testing.T) {
	test(t, []int{3, 4, -1, 1}, 2)
}
