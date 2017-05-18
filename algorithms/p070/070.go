package p070

/**
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.
 */

// 斐波那契数列
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	n_2, n_1 := 1, 2
	for n > 2 {
		n_1, n_2 = n_1+n_2, n_1
		n--
	}
	return n_1
}
