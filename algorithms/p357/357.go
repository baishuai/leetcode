package p357

/**
Given a non-negative integer n, count all numbers with unique digits, x, where 0 ≤ x < 10n.

Example:
Given n = 2, return 91. (The answer should be the total numbers in the range of 0 ≤ x < 100, excluding [11,22,33,44,55,66,77,88,99])
**/

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	n--
	res := 10
	cnt := 9
	unique := 9
	for n > 0 && unique > 0 {
		res += cnt * unique
		cnt *= unique
		n--
		unique--
	}
	return res
}
