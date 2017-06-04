package p304

/**
Given a 2D matrix matrix, find the sum of the elements inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).

Range Sum Query 2D
The above rectangle (with the red border) is defined by (row1, col1) = (2, 1) and (row2, col2) = (4, 3), which contains sum = 8.

Example:
Given matrix = [
  [3, 0, 1, 4, 2],
  [5, 6, 3, 2, 1],
  [1, 2, 0, 1, 5],
  [4, 1, 0, 1, 7],
  [1, 0, 3, 0, 5]
]

sumRegion(2, 1, 4, 3) -> 8
sumRegion(1, 1, 2, 2) -> 11
sumRegion(1, 2, 2, 4) -> 12


You may assume that the matrix does not change.
There are many calls to sumRegion function.
You may assume that row1 ≤ row2 and col1 ≤ col2.

**/

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil {
		return NumMatrix{sums: [][]int{{0}}}
	}
	sums := make([][]int, len(matrix)+1)
	for i := 0; i < len(sums); i++ {
		sums[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i < len(sums); i++ {
		copy(sums[i][1:], matrix[i-1])
		for j := 1; j < len(sums[i]); j++ {
			sums[i][j] += sums[i][j-1]
		}
	}

	for j := 1; j < len(sums[0]); j++ {
		for i := 1; i < len(sums); i++ {
			sums[i][j] += sums[i-1][j]
		}
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sums[row2+1][col2+1] - this.sums[row1][col2+1] -
		this.sums[row2+1][col1] + this.sums[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
