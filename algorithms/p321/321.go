package p321

/**
Given two arrays of length m and n with digits 0-9 representing two numbers.
Create the maximum number of length k <= m + n from digits of the two.
The relative order of the digits from the same array must be preserved.
Return an array of the k digits. You should try to optimize your time and space complexity.

Example 1:
nums1 = [3, 4, 6, 5]
nums2 = [9, 1, 2, 5, 8, 3]
k = 5
return [9, 8, 6, 5, 3]

Example 2:
nums1 = [6, 7]
nums2 = [6, 0, 4]
k = 5
return [6, 7, 6, 0, 4]

Example 3:
nums1 = [3, 9]
nums2 = [8, 9]
k = 3
return [9, 8, 9]
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ans := make([]int, k)
	n, m := len(nums1), len(nums2)
	for i := max(0, k-m); i <= k && i <= n; i++ {
		candidate := merge(maxOneArray(nums1, i), maxOneArray(nums2, k-i), k)
		if greater(candidate, 0, ans, 0) {
			ans = candidate
		}
	}
	return ans
}

func merge(nums1, nums2 []int, k int) []int {
	ans := make([]int, k)
	for i, j, r := 0, 0, 0; r < k; r++ {
		if greater(nums1, i, nums2, j) {
			ans[r] = nums1[i]
			i++
		} else {
			ans[r] = nums2[j]
			j++
		}
	}
	return ans
}

func greater(nums1 []int, i int, nums2 []int, j int) bool {
	for i < len(nums1) && j < len(nums2) && nums1[i] == nums2[j] {
		i++
		j++
	}
	return j == len(nums2) || (i < len(nums1) && nums1[i] > nums2[j])
}

func maxOneArray(nums []int, k int) []int {
	ans := make([]int, k)
	n := len(nums)
	i := 0
	for j := 0; j < k; j++ {
		largest := 0
		ix := i
		for i < n-k+1+j {
			if nums[i] > largest {
				largest = nums[i]
				ix = i
			}
			i++
		}
		i = ix + 1
		ans[j] = largest
	}
	return ans
}
