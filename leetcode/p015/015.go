package p015

import "sort"

/**
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */

func threeSum(nums []int) [][]int {
	ansSet := make([][]int, 0)
	sort.Ints(nums)
	for first := 0; first < len(nums)-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := 0 - nums[first]
		second := first + 1
		third := len(nums) - 1
		for second < third {
			curSum := nums[second] + nums[third]
			if curSum == target { // one ans
				ans := []int{nums[first], nums[second], nums[third]}
				ansSet = append(ansSet, ans)
				for second < third && nums[second] == ans[1] {
					second++
				}
				for third > second && nums[third] == ans[2] {
					third--
				}
			} else if curSum < target {
				second++
			} else {
				third--
			}
		}
	}
	return ansSet
}
