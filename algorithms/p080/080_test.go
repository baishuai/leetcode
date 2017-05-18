package p080

import "testing"

func test(t *testing.T, nums []int, exp []int) {
	if removeDuplicates(nums) != len(exp) {
		t.Fatal("Error answer length")
	}
	for i, v := range exp {
		if nums[i] != v {
			t.Error("error answer value", i, v, nums[i])
		}
	}
}

func TestExample(t *testing.T) {
	test(t, []int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3})
}
