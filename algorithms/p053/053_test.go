package p053

import "testing"

func TestExample(t *testing.T) {
	if maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}) != 6 {
		t.Fatal("error answer!")
	}
}