package p118

/**
Given numRows, generate the first numRows of Pascal's triangle.

For example, given numRows = 5,
Return

[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	if numRows >= 1 {
		res[0] = append(res[0], 1)
	}
	for i := 1; i < numRows; i++ {
		res[i] = append(res[i], 1)
		for j := 0; j < i-1; j++ {
			res[i] = append(res[i], res[i-1][j]+res[i-1][j+1])
		}
		res[i] = append(res[i], 1)
	}
	return res
}
