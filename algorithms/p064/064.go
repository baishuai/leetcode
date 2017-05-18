package p064

/**
Given a m x n grid filled with non-negative numbers, find a path
 from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.
 */

func minPathSum(grid [][]int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}
