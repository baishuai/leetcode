package p048

/**
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Follow up:
Could you do this in-place?
 */

func rotate(matrix [][]int) {
	i, j, n := 0, len(matrix)-1, len(matrix)
	for j > i {
		for k := 0; k < n; k++ {
			matrix[i][k], matrix[j][k] = matrix[j][k], matrix[i][k]
		}
		i++
		j--
	}

	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
