package p221

/**
Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

For example, given the following matrix:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0
Return 4.
**/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximalSquare(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rst := 0
	dp, preDp := make([]int, len(matrix[0])+1), make([]int, len(matrix[0])+1)

	for _, line := range matrix {
		dp, preDp = preDp, dp
		for i, b := range line {
			if b == '0' {
				dp[i+1] = 0
			} else {
				dp[i+1] = min(dp[i], min(preDp[i], preDp[i+1])) + 1
			}
			rst = max(rst, dp[i+1]*dp[i+1])
		}
	}
	return rst
}
