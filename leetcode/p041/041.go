package p041

/**
Given an unsorted integer array, find the first missing positive integer.

For example,
Given [1,2,0] return 3,
and [3,4,-1,1] return 2.

Your algorithm should run in O(n) time and uses constant space.
 */

func firstMissingPositive(nums []int) int {

	for i, v := range nums {
		if v < 0 {
			nums[i] = 0
		}
	}
	flag, lens := 1<<(32-1), len(nums)
	mask, ans := flag-1, lens+1
	for _, v := range nums {
		actual := v & mask
		if actual > 0 && actual <= lens {
			nums[actual-1] = nums[actual-1] | flag
		}
	}

	for i, v := range nums {
		if (v & flag) == 0 {
			ans = i + 1
			break
		}
	}
	return ans
}
