package p077

/**
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

For example,
If n = 4 and k = 2, a solution is:

[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	if k == 0 {
		return res
	}
	one := make([]int, k)
	count := 1
	for i := 0; i < k; i++ {
		one[i] = i + 1
		count *= n - i
	}
	for i := 1; i <= k; i++ {
		count /= i
	}

	for count > 0 {
		onecase := make([]int, k)
		copy(onecase, one)
		res = append(res, onecase)
		for i := k - 1; i >= 0; i-- {
			if one[i] < (n - k + 1 + i) {
				one[i] += 1
				for j := i + 1; j < k; j++ {
					one[j] = one[j-1] + 1
				}
				break
			}
		}
		count--
	}
	return res
}
