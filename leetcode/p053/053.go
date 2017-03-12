package p053

/**
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.
 */

func maxSubArray(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}
	max := nums[0]
	p := 0
	for _, v := range nums {
		if v+p > v {
			p = v + p
		} else {
			p = v
		}
		if p > max {
			max = p
		}
	}
	return max
}
