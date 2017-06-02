package p279

/**
Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

For example, given n = 12, return 3 because 12 = 4 + 4 + 4; given n = 13, return 2 because 13 = 4 + 9.

**/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	nums := make([]int, n+1)

	for i := 1; i <= n; i++ {
		ns := i
		for k := 1; i-k*k >= 0; k++ {
			ns = min(nums[i-k*k]+1, ns)
		}
		nums[i] = ns
	}
	return nums[n]
}
