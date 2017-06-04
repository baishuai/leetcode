package p088

/**
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional
 elements from nums2. The number of elements initialized in nums1 and nums2 are m and n respectively.
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := nums1
	nums1 = nums1[len(nums1)-m:]
	copy(nums1, nums[:m])
	l, r := 0, 0
	for l < m && r < n {
		if nums1[l] < nums2[r] {
			nums[l+r] = nums1[l]
			l++
		} else {
			nums[l+r] = nums2[r]
			r++
		}
	}
	if l < m {
		copy(nums[l+r:], nums1[l:])
	} else if r < n {
		copy(nums[l+r:], nums2[r:])
	}
}
