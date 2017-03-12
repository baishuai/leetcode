package p075

import "testing"

func test(t *testing.T, nums []int, exp []int) {
	sortColors(nums)
	if len(nums) != len(exp) {
		t.Fatal("Error answer length")
	}
	for i, v := range exp {
		if nums[i] != v {
			t.Error("error value", i, nums[i], v)
		}
	}
}

func TestEXample0(t *testing.T) {
	test(t, []int{0, 1, 0, 1, 2, 1, 2, 0, 2, 1, 0, 0, 1, 2}, []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2})
}

func TestEXample1(t *testing.T) {
	test(t, []int{0, 0}, []int{0, 0})
}

func TestEXample2(t *testing.T) {
	test(t, []int{2, 2, 2}, []int{2, 2, 2})
}

func TestEXample3(t *testing.T) {
	test(t, []int{1, 1, 1}, []int{1, 1, 1})
}

func TestEXample4(t *testing.T) {
	test(t, []int{1, 1, 1, 0}, []int{0, 1, 1, 1})
}

func TestEXample5(t *testing.T) {
	test(t, []int{1, 1, 1, 2}, []int{1, 1, 1, 2})
}
