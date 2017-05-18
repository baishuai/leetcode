package p011

import "testing"

func TestTwo(t *testing.T) {
	two := []int{2, 3}
	result := maxArea(two)
	if result != 2 {
		t.Fatal("error answer", result)
	}
}

func TestThree(t *testing.T) {
	two := []int{2, 4, 3}
	result := maxArea(two)
	if result != 4 {
		t.Fatal("error answer", result)
	}
}
