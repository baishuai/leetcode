package p016

import (
	"sort"
)

/**
Given an array S of n integers, find three integers in S such that
 the sum is closest to a given number, target.
 Return the sum of the three integers.
 You may assume that each input would have exactly one solution.

    For example, given array S = {-1 2 1 -4}, and target = 1.

    The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 */

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	lens := len(nums)
	if lens < 3 {
		return 0
	}
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for first := 0; first < lens-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second := first + 1
		third := lens - 1
		for second < third {
			curSum := nums[first] + nums[second] + nums[third]
			if curSum == target {
				return curSum
			}
			if abs(target-curSum) < abs(target-closest) {
				closest = curSum
			}
			if curSum < target {
				second++
				for second < lens && nums[second] == nums[second-1] {
					second++
				}
			} else {
				third--
				for third > 0 && nums[third] == nums[third+1] {
					third--
				}
			}
		}
	}
	return closest
}
