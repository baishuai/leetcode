package p096

/**
Given n, how many structurally unique BST's (binary search trees)
 that store values 1...n?

For example,
Given n = 3, there are a total of 5 unique BST's.
 */

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}
	nums := make([]int, n+1)
	nums[0], nums[1] = 1, 1
	for i := 2; i <= n; i++ {
		l, h := 0, i-1
		sum := 0
		for l < h {
			sum += nums[l] * nums[h]
			l++
			h--
		}
		sum = sum * 2
		if l == h {
			sum += nums[l] * nums[l]
		}
		nums[i] = sum
	}
	return nums[n]
}
