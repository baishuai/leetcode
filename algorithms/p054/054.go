package p054

/**
Given a matrix of m x n elements (m rows, n columns),
 return all elements of the matrix in spiral order.

For example,
Given the following matrix:

[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
You should return [1,2,3,6,9,8,7,4,5].
*/

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	res := make([]int, m*n)

	cnt := 0
	total := m * n
	i, j, dest := 0, 0, 0
	for cnt < total {
		for dest = j + n; j < dest; j++ {
			res[cnt] = matrix[i][j]
			cnt++
		}
		if cnt >= total {
			break
		}
		j--
		i++
		for dest = i + (m - 1); i < dest; i++ {
			res[cnt] = matrix[i][j]
			cnt++
		}
		if cnt >= total {
			break
		}
		i--
		j--
		for dest = j - (n - 1); j > dest; j-- {
			res[cnt] = matrix[i][j]
			cnt++
		}
		if cnt >= total {
			break
		}
		j++
		i--
		for dest = i - (m - 2); i > dest; i-- {
			res[cnt] = matrix[i][j]
			cnt++
		}
		i++
		j++
		m -= 2
		n -= 2
	}
	return res
}
