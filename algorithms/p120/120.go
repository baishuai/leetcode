package p120

/**
Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

 */
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumTotal(triangle [][]int) int {
	size := len(triangle)
	if size == 0 {
		return 0
	}
	dpPre := make([]int, size)
	dpNow := make([]int, size)
	dpNow[0] = triangle[0][0]
	for i := 1; i < size; i++ {
		row := triangle[i]
		dpPre, dpNow = dpNow, dpPre
		for j := 0; j <= i; j++ {
			if j == 0 {
				dpNow[j] = dpPre[j] + row[j]
			} else if j == i {
				dpNow[j] = dpPre[j-1] + row[j]
			} else {
				dpNow[j] = min(dpPre[j-1], dpPre[j]) + row[j]
			}
		}
	}
	minValue := dpNow[0]
	for i := 1; i < size; i++ {
		if dpNow[i] < minValue {
			minValue = dpNow[i]
		}
	}
	return minValue
}
