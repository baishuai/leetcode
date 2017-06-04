package p215

/**
Find the kth largest element in an unsorted array. Note that it is the kth largest element
 in the sorted order, not the kth distinct element.

For example,
Given [3,2,1,5,6,4] and k = 2, return 5.

Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/

func findKthLargest(nums []int, k int) int {
	v := nums[0]
	i, j := 0, len(nums)-1
	for i < j {
		for i < len(nums) && nums[i] >= v {
			i++
		}
		for j > 0 && nums[j] < v {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[0], nums[j] = nums[j], nums[0]
	if k == j+1 {
		return nums[j]
	} else if k < j+1 {
		return findKthLargest(nums[:j], k)
	} else {
		return findKthLargest(nums[j+1:], k-j-1)
	}
}
