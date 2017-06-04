package p063

/**
Follow up for "Unique Paths":

Now consider if some obstacles are added to the grids.
 How many unique paths would there be?

An obstacle and empty space is marked as 1 and 0 respectively in the grid.

For example,
There is one obstacle in the middle of a 3x3 grid as illustrated below.

[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
The total number of unique paths is 2.
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	obstacleGrid[0][0] = 1 - obstacleGrid[0][0]
	if obstacleGrid[0][0] == 0 {
		return 0
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] != 1 {
			obstacleGrid[0][j] = obstacleGrid[0][j-1]
		} else {
			obstacleGrid[0][j] = 0
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] != 1 {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		} else {
			obstacleGrid[i][0] = 0
		}

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}
