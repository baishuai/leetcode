package p461

/**
The Hamming distance between two integers is the number of positions
 at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 2^31.
 */

func hammingDistance(x int, y int) int {
	xor := x ^ y
	ans := 0
	for xor != 0 {
		ans += xor & 1
		xor = xor >> 1
	}
	return ans
}
