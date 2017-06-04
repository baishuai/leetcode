package p172

/**
Given an integer n, return the number of trailing zeroes in n!.

Note: Your solution should be in logarithmic time complexity.
*/

func trailingZeroes(n int) int {
	cnt := 0
	for i := 5; i <= n; i *= 5 {
		cnt += n / i
	}
	return cnt
}
