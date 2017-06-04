package p119

/**
Given an index k, return the kth row of the Pascal's triangle.

For example, given k = 3,
Return [1,3,3,1].
*/

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		res[i] = 1
		for j := i - 1; j > 0; j-- {
			res[j] = res[j] + res[j-1]
		}
		res[0] = 1
	}
	return res
}
