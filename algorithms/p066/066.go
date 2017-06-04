package p066

/**
Given a non-negative integer represented as a non-empty array of digits,
 plus one to the integer.

You may assume the integer do not contain any leading zero,
 except the number 0 itself.

The digits are stored such that the most significant digit is at the head of the list.
*/

func plusOne(digits []int) []int {
	ans := make([]int, len(digits)+1)
	copy(ans[1:], digits)
	for i := len(ans) - 1; i >= 0; i-- {
		ans[i] += 1
		if ans[i] <= 9 {
			break
		} else {
			ans[i] = 0
		}
	}
	if ans[0] == 0 {
		ans = ans[1:]
	}
	return ans
}
