package p326

/**
Given an integer, write a function to determine if it is a power of three.

Follow up:
Could you do it without using any loop / recursion?
*/

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
