package p695

func dfs(grid [][]int, n, m int, i, j int) int {
	if i < 0 || i >= n || j < 0 || j >= m {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + dfs(grid, n, m, i-1, j) + dfs(grid, n, m, i+1, j) +
		dfs(grid, n, m, i, j-1) + dfs(grid, n, m, i, j+1)
}

func maxAreaOfIsland(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	maxArea := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				v := dfs(grid, n, m, i, j)
				if v > maxArea {
					maxArea = v
				}
			}
		}
	}
	return maxArea
}
