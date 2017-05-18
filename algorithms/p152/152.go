package p152

/**
Find the contiguous subarray within an array (containing at least one number)
 which has the largest product.

For example, given the array [2,3,-2,4],
the contiguous subarray [2,3] has the largest product = 6.
 */

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func maxProduct(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}
	maxp, minp := nums[0], nums[0]
	res := nums[0]
	for _, v := range nums[1:] {
		maxp, minp = max(maxp*v, minp*v, v), min(maxp*v, minp*v, v)
		if maxp > res{
			res = maxp
		}
	}
	return res
}
