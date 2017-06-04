package p074

/**
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
For example,

Consider the following matrix:

[
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	l, h := 0, m*n
	for l < h {
		mid := l + ((h - l) >> 1)
		if matrix[mid/n][mid%n] < target {
			l = mid + 1
		} else {
			h = mid
		}
	}
	return (l < m*n) && (matrix[l/n][l%n] == target)
}
