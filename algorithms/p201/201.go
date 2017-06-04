package p201

/**
Given a range [m, n] where 0 <= m <= n <= 2147483647,
 return the bitwise AND of all numbers in this range, inclusive.

For example, given the range [5, 7], you should return 4.
*/

func rangeBitwiseAnd(m int, n int) int {
	factor := 1
	for m != n {
		m >>= 1
		n >>= 1
		factor <<= 1
	}
	return factor * m
}
