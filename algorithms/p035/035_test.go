package p035

import "testing"

func test(t *testing.T, nums []int, target, exp int) {
	ans := searchInsert(nums, target)
	if ans != exp {
		t.Fatal("error answer", exp, "get", ans)
	}
}

func TestExample1(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	test(t, nums, 5, 2)
}

func TestExample2(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	test(t, nums, 2, 1)
}

func TestExample3(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	test(t, nums, 7, 4)
}

func TestExample4(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	test(t, nums, 0, 0)
}

func TestExample5(t *testing.T) {
	nums := []int{0}
	test(t, nums, 0, 0)
}
