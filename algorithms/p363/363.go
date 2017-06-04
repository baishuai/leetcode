package p363

import (
	"math"
)

//363. Max Sum of Rectangle No Larger Than K

/**
Given a non-empty 2D matrix matrix and an integer k, find the max sum of a rectangle in the matrix such that its sum is no larger than k.

Example:
Given matrix = [
  [1,  0, 1],
  [0, -2, 3]
]
k = 2
The answer is 2. Because the sum of rectangle [[0, 1], [-2, 3]] is 2 and 2 is the max number no larger than k (k = 2).

Note:
The rectangle inside the matrix must have an area > 0.
What if the number of rows is much larger than the number of columns?
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 {
		return 0
	}
	rows := len(matrix)
	cols := len(matrix[0])
	res := math.MinInt32
	sums := make([]int, rows)
	zeros := make([]int, rows)
	sums2d := make([]int, rows+1)
	helper := make([]int, rows+1)
	for l := 0; l < cols; l++ {
		copy(sums, zeros)
		for r := l; r < cols; r++ {
			sums2d[0] = 0
			for i := 0; i < rows; i++ {
				sums[i] += matrix[i][r]
				sums2d[i+1] = sums2d[i] + sums[i]
			}
			ttm := mergeSort(sums2d, helper, k)
			res = max(res, ttm)

		}
	}
	return res
}

func mergeSort(sum, helper []int, k int) int {
	if len(sum) <= 1 {
		return math.MinInt32
	}
	mid := len(sum) / 2
	ans := mergeSort(sum[:mid], helper, k)
	if ans == k {
		return k
	}
	ans = max(ans, mergeSort(sum[mid:], helper, k))
	if ans == k {
		return k
	}
	j, p := mid, mid
	r := 0
	//copy(helper, sum[:mid])
	for i := 0; i < mid; i++ {
		for j < len(sum) && sum[j]-sum[i] <= k {
			j++
		}
		if j-1 >= mid {
			ans = max(ans, sum[j-1]-sum[i])
			if ans == k {
				return k
			}
		}
		for p < len(sum) && sum[p] < sum[i] {
			helper[r] = sum[p]
			r++
			p++
		}
		helper[r] = sum[i]
		r++
	}
	copy(sum, helper[:r])
	return ans
}
