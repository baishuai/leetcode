package p059

/**
Given an integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

For example,
Given n = 3,

You should return the following matrix:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
 */

func generateMatrix(n int) [][]int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}

	rowStart := 0
	rowEnd := n - 1
	colStart := 0
	colEnd := n - 1
	fill := 1

	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart; i <= colEnd; i++ {
			mat[rowStart][i] = fill
			fill++
		}
		rowStart ++

		for i := rowStart; i <= rowEnd; i++ {
			mat[i][colEnd] = fill
			fill++
		}
		colEnd--

		for i := colEnd; i >= colStart; i-- {
			mat[rowEnd][i] = fill
			fill++
		}
		rowEnd--

		for i := rowEnd; i >= rowStart; i-- {
			mat[i][colStart] = fill
			fill++
		}
		colStart++
	}
	return mat
}
