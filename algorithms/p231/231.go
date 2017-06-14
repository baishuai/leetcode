package p231

/**
Given an integer, write a function to determine if it is a power of two.
*/

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	c := 1
	for x := n &^ c; x != 0; c <<= 1 {
		x = n &^ c
		if x < n && x > 0 {
			return false
		}
	}
	return true
}
