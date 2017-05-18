package p312

/**
Given n balloons, indexed from 0 to n-1. Each balloon is painted with a number
 on it represented by array nums. You are asked to burst all the balloons.
 If the you burst balloon i you will get nums[left] * nums[i] * nums[right] coins.
 Here left and right are adjacent indices of i. After the burst, the left and right then becomes adjacent.

Find the maximum coins you can collect by bursting the balloons wisely.

Note:
(1) You may imagine nums[-1] = nums[n] = 1. They are not real therefore you can not burst them.
(2) 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100

Example:

Given [3, 1, 5, 8]

Return 167

    nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
   coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxCoins(nums []int) int {

	balloons := make([]int, len(nums)+2)
	balloons[0] = 1
	copy(balloons[1:], nums)
	balloons[len(balloons)-1] = 1
	sum := 0
	if nums != nil && len(nums) > 0 {

		memo := make([][]int, len(balloons))
		for i := 0; i < len(nums); i++ {
			memo[i] = make([]int, len(balloons))
		}

		sum += burst(memo, balloons, 0, len(balloons)-1)
	}
	return sum
}

func burst(memo [][]int, nums []int, left, right int) int {
	if left+1 == right {
		return 0
	}
	if memo[left][right] > 0 {
		return memo[left][right]
	}

	ans := 0
	for i := left + 1; i < right; i++ {
		ans = max(ans, nums[left]*nums[i]*nums[right]+
			burst(memo, nums, left, i)+ burst(memo, nums, i, right))
	}
	memo[left][right] = ans
	return ans
}
