package p004

import (
	"testing"
)

func TestSame(t *testing.T) {
	nums1 := []int{2, 2}
	nums2 := []int{2, 2}

	mean := findMedianSortedArrays(nums1, nums2)

	if mean != 2 {
		t.Fatal("error answer")
	}
}

func TestMean(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	mean := findMedianSortedArrays(nums1, nums2)

	if mean != 2 {
		t.Fatal("error answer")
	}
}

func Test0(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{1, 2, 3}

	mean := findMedianSortedArrays(nums1, nums2)

	if mean != 2 {
		t.Fatal("error answer")
	}
}

func Test1(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{1, 2, 3}

	mean := findMedianSortedArrays(nums1, nums2)

	if mean != 1.5 {
		t.Fatal("error answer")
	}
}

func Test2(t *testing.T) {
	nums2 := []int{3}
	nums1 := []int{1, 2, 3}

	mean := findMedianSortedArrays(nums1, nums2)
	if mean != 2.5 {
		t.Fatal("error answer")
	}
}

func Test3(t *testing.T) {
	nums2 := []int{3}
	nums1 := []int{1, 2, 3, 4, 5}

	mean := findMedianSortedArrays(nums1, nums2)
	if mean != 3 {
		t.Fatal("error answer")
	}
}

func Test4(t *testing.T) {
	nums2 := []int{2}
	nums1 := []int{1, 2, 3, 4}

	mean := findMedianSortedArrays(nums1, nums2)
	if mean != 2 {
		t.Fatal("error answer")
	}
}
