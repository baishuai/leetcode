package p329

/**
Given an integer matrix, find the length of the longest increasing path.

From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).

Example 1:

nums = [
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
Return 4
The longest increasing path is [1, 2, 6, 9].

Example 2:

nums = [
  [3,4,5],
  [3,2,6],
  [2,2,1]
]
Return 4
The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
 */

func longestIncreasingPath(matrix [][]int) int {

	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	mem := make([][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = -1
		}
	}

	longest := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			longest = max(longest, dfs(matrix, mem, i, j))
		}
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(matrix [][]int, meme [][]int, i, j int) int {
	if meme[i][j] >= 0 {
		return meme[i][j]
	}

	longest := 1
	if i > 0 && matrix[i-1][j] < matrix[i][j] {
		longest = max(longest, 1+dfs(matrix, meme, i-1, j))
	}
	if j > 0 && matrix[i][j-1] < matrix[i][j] {
		longest = max(longest, 1+dfs(matrix, meme, i, j-1))
	}
	if i < len(matrix)-1 && matrix[i+1][j] < matrix[i][j] {
		longest = max(longest, 1+dfs(matrix, meme, i+1, j))
	}
	if j < len(matrix[0])-1 && matrix[i][j+1] < matrix[i][j] {
		longest = max(longest, 1+dfs(matrix, meme, i, j+1))
	}
	meme[i][j] = longest
	return longest
}
