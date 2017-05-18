package p016

import "testing"

func TestExample(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	if ans := threeSumClosest(nums, 1); ans != 2 {
		t.Fatal("error answer", ans)
	}
}

func TestNoless(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	if ans := threeSumClosest(nums, 1); ans != 6 {
		t.Fatal("error answer", ans)
	}
}

func TestNoGreater(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	if ans := threeSumClosest(nums, 10); ans != 9 {
		t.Fatal("error answer", ans)
	}
}
