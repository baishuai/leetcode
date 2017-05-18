package p034

import "testing"

func test(t *testing.T, nums []int, target int, expl int, expr int) {
	ans := searchRange(nums, target)
	if len(ans) != 2 {
		t.Fatal("errow answer length", ans)
	}
	if ans[0] != expl || ans[1] != expr {
		t.Fatal("error answer val", ans)
	}
}

func TestExample1(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	test(t, nums, 8, 3, 4)
}

func TestExample2(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	test(t, nums, 7, 1, 2)
}

func TestExtra1(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	test(t, nums, 5, 0, 0)
}

func TestExtra2(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10, 10}
	test(t, nums, 10, 5, 6)
}

func TestExtra3(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	test(t, nums, 9, -1, -1)
}

func TestExtra4(t *testing.T) {
	nums := []int{5}
	test(t, nums, 9, -1, -1)
}

func TestExtra5(t *testing.T) {
	nums := []int{}
	test(t, nums, 9, -1, -1)
}

func TestExtra6(t *testing.T) {
	nums := []int{5}
	test(t, nums, 5, 0, 0)
}

func TestExtra7(t *testing.T) {
	nums := []int{5, 6, 7, 8}
	test(t, nums, 7, 2, 2)
}
