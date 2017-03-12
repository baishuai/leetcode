package p026

import "testing"

func TestExample(t *testing.T) {
	nums := []int{1, 1, 2}
	ans := []int{1, 2}
	if length := removeDuplicates(nums); length != 2 {
		t.Fatal("error length")
		for i, v := range ans {
			if nums[i] != v {
				t.Fatal("error answer")
			}
		}
	}
}
