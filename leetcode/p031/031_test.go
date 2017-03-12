package p031

import "testing"

func test(t *testing.T, nums []int, exp []int) {
	nextPermutation(nums)
	if len(nums) != len(exp) {
		t.Fatal("error answer length", len(nums))
	}
	for i, v := range exp {
		if v != nums[i] {
			t.Error("error answer val", nums[i], "expected:", v)
		}
	}
}

func TestExample1(t *testing.T) {
	nums := []int{1, 2, 3}
	exp := []int{1, 3, 2}
	test(t, nums, exp)
}

func TestExample2(t *testing.T) {
	nums := []int{3, 2, 1}
	exp := []int{1, 2, 3}
	test(t, nums, exp)
}

func TestExample3(t *testing.T) {
	nums := []int{1, 1, 5}
	exp := []int{1, 5, 1}
	test(t, nums, exp)
}

func TestExample4(t *testing.T) {
	nums := []int{1,2,5,3,6,4,2}
	exp := []int{1,2,5,4,2,3,6}
	test(t, nums, exp)
}

func TestExample5(t *testing.T) {
	nums := []int{1,5,1}
	exp := []int{5,1,1}
	test(t, nums, exp)
}
