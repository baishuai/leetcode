package p073

/**
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in place.

Follow up:
Did you use extra space?
A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	fr, fc := false, false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				fr = fr || (i == 0)
				fc = fc || (j == 0)
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if fr {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if fc {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
