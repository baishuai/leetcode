package p001

import "testing"

func TestSumTwo(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	if len(result) != 2 {
		t.Fatal("result length error")
	}
	if result[0] != 0 || result[1] != 1 {
		t.Fatal("get error result")
	}
}
