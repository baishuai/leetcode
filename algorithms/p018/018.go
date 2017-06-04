package p018

import (
	"sort"
)

/**
Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note: The solution set must not contain duplicate quadruplets.

For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func fourSum(nums []int, target int) [][]int {

	ansSet := make([][]int, 0)
	sort.Ints(nums)
	var lastAns []int
	for head := 0; head < len(nums)-3; head++ {
		if lastAns != nil && nums[head] == lastAns[0] {
			continue
		}
		lastAns = nil
		tmh := target - nums[head]
		for first := head + 1; first < len(nums)-2; first++ {
			if lastAns != nil && nums[first] == lastAns[1] {
				continue
			}
			tmhs := tmh - nums[first]
			second := first + 1
			third := len(nums) - 1
			for second < third {
				curSum := nums[second] + nums[third]
				if curSum == tmhs { // one ans
					lastAns = []int{nums[head], nums[first], nums[second], nums[third]}
					ansSet = append(ansSet, lastAns)
					for second < third && nums[second] == lastAns[2] {
						second++
					}
					for third > second && nums[third] == lastAns[3] {
						third--
					}
				} else if curSum < tmhs {
					second++
				} else {
					third--
				}
			}
		}
	}
	return ansSet
}
