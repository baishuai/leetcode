package p202

/**
Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Example: 19 is a happy number

12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
 */

func digitSquare(m int) int {
	sum, tmp := 0, 0
	for m > 0 {
		tmp = m % 10
		sum += tmp * tmp
		m = m / 10
	}
	return sum
}

func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = digitSquare(slow)
		fast = digitSquare(fast)
		fast = digitSquare(fast)
		if slow == fast {
			break
		}
	}
	return slow == 1
}
