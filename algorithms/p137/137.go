package p137

/**
Given an array of integers, every element appears three times except for one,
 which appears exactly once. Find that single one.

Note:
Your algorithm should have a linear runtime complexity.
 Could you implement it without using extra memory?
 */

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}
