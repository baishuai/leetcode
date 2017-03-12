package p045

import "testing"

func test(t *testing.T, nums []int, exp int) {
	if ans := jump(nums); ans != exp {
		t.Fatal("error answer", ans, "expected", exp)
	}
}

func TestExample(t *testing.T) {
	test(t, []int{2, 3, 1, 1, 4}, 2)
}
