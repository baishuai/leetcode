package p343

/**
Given a positive integer n, break it into the sum of at least two positive integers and maximize the product of those integers. Return the maximum product you can get.

For example, given n = 2, return 1 (2 = 1 + 1); given n = 10, return 36 (10 = 3 + 3 + 4).

Note: You may assume that n is not less than 2 and not larger than 58.
**/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak(n int) int {

	dp := make([]int, n+1)
	dp[1] = 1
	max := 0
	for i := 2; i <= n; i++ {
		max = 0
		for l := 1; l <= i/2; l++ {
			max = Max(max, dp[l]*dp[i-l])
		}
		if i == n {
			break
		}
		dp[i] = Max(max, i)
	}
	return max
}
