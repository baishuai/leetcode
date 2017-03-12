package p004

import (
	"testing"
)

func TestSame(t *testing.T)  {
	nums1 := []int{2, 2}
	nums2 := []int{2,2}

	mean := findMedianSortedArrays(nums1, nums2)

	if mean != 2{
		t.Fatal("error answer")
	}
}

func TestMean(t *testing.T)  {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	mean := findMedianSortedArrays(nums1, nums2)

	if mean != 2{
		t.Fatal("error answer")
	}
}