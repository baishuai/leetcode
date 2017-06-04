package p576

/**
There is an m by n grid with a ball. Given the start coordinate (i,j) of the ball,
you can move the ball to adjacent cell or cross the grid boundary in four directions
 (up, down, left, right).
However, you can at most move N times.
Find out the number of paths to move the ball out of grid boundary.
The answer may be very large, return it after mod 109 + 7.

Note:
Once you move the ball out of boundary, you cannot move it back.
The length and height of the grid is in range [1,50].
N is in range [0,50].
*/

//从点（i，j）走到边界可以反过来考虑，从任意边界走到点（i，j）

func findPaths(m int, n int, N int, ii int, jj int) int {
	dp := make([][]int, m)
	preDp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		preDp[i] = make([]int, n)
	}

	for N > 0 {
		N--
		dp, preDp = preDp, dp
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				dp[i][j] = 0
				if i > 0 {
					dp[i][j] += preDp[i-1][j]
				} else {
					dp[i][j] += 1
				}
				if i < m-1 {
					dp[i][j] += preDp[i+1][j]
				} else {
					dp[i][j] += 1
				}
				if j > 0 {
					dp[i][j] += preDp[i][j-1]
				} else {
					dp[i][j] += 1
				}
				if j < n-1 {
					dp[i][j] += preDp[i][j+1]
				} else {
					dp[i][j] += 1
				}
				dp[i][j] = dp[i][j] % 1000000007
			}
		}
	}
	return dp[ii][jj]
}
