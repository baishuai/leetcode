package p007

/**
Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

click to show spoilers.

Have you thought about this?
Here are some good questions to ask before coding. Bonus points for you if you have already thought through this!

If the integer's last digit is 0, what should the output be? ie, cases such as 10, 100.

Did you notice that the reversed integer might overflow? Assume the input is a 32-bit integer, then the reverse of 1000000003 overflows. How should you handle such cases?

For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.


 */

func reverse(x int) int {
	result := 0

	flag := x < 0
	if flag {
		x = -x
	}
	for x > 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result < 0 || result > (1<<31-1) {
		result = 0
	}
	if flag {
		result = -result
	}
	return result
}
