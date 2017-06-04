package p233

/**
Given an integer n, count the total number of digit 1 appearing in all non-negative
 integers less than or equal to n.

For example:
Given n = 13,
Return 6, because digit 1 occurred in the following numbers: 1, 10, 11, 12, 13.
*/

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	} else if n < 10 {
		return 1
	}
	prod, pow := 1, 0
	for n/prod >= 10 {
		prod *= 10
		pow++
	}
	ones := pow * prod / 10
	ones = ones * (n / prod)
	if n/prod >= 2 {
		ones += prod
	} else {
		ones += n - prod + 1
	}
	return ones + countDigitOne(n%prod)
}
