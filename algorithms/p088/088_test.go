package p088

import "testing"

func test(t *testing.T, nums1 []int, m int, nums2 []int, n int, exp []int) {
	merge(nums1, m, nums2, n)
	for i, v := range exp {
		if nums1[i] != v {
			t.Error("error answer", i, nums1[i])
		}
	}
}

func TestExample0(t *testing.T) {
	test(t, []int{0}, 0, []int{1}, 1, []int{1})
}

func TestExample1(t *testing.T) {
	test(t, []int{2, 0}, 1, []int{1}, 1, []int{1, 2})
}
