package p213

/**
Note: This is an extension of House Robber.

After robbing those houses on that street, the thief has found himself a new place for his thievery
 so that he will not get too much attention. This time, all houses at this place are arranged in a circle.
 That means the first house is the neighbor of the last one. Meanwhile, the security system for these houses
 remain the same as for those in the previous street.

Given a list of non-negative integers representing the amount of money of each house,
 determine the maximum amount of money you can rob tonight without alerting the police.
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob_inner(nums []int) int {
	//[1.2.3.4.5]
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(rob_inner(nums[1:]), rob_inner(nums[:len(nums)-1]))
}
