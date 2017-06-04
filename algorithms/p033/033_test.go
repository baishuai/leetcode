package p033

import "testing"

func test(t *testing.T, nums []int, target int, exp int) {
	ans := search(nums, target)
	if ans != exp {
		t.Fatal("errow answer", ans)
	}
}

func TestExample(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	test(t, nums, 5, 1)
}

func TestExtra1(t *testing.T) {
	nums := []int{4, 2}
	test(t, nums, 5, -1)
}

func TestExtra2(t *testing.T) {
	nums := []int{4, 2}
	test(t, nums, 2, 1)
}

func TestExtra3(t *testing.T) {
	nums := []int{4, 2}
	test(t, nums, 4, 0)
}

func TestExtra4(t *testing.T) {
	nums := []int{4}
	test(t, nums, 4, 0)
}

func TestExtra5(t *testing.T) {
	nums := []int{4}
	test(t, nums, 5, -1)
}
