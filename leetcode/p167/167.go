package p167

import (
	"sort"
)

/**
Given an array of integers that is already sorted in ascending order,
 find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that
 they add up to the target, where index1 must be less than index2.
  Please note that your returned answers (both index1 and index2) are not zero-based.

You may assume that each input would have exactly
 one solution and you may not use the same element twice.

Input: numbers={2, 7, 11, 15}, target=9
Output: index1=1, index2=2
 */

func twoSum(numbers []int, target int) []int {

	i, j := 0, len(numbers)-1
	for {
		if numbers[i]+numbers[j] > target {
			j = sort.SearchInts(numbers[i:], target-numbers[i]) + i
			if j == len(numbers) || numbers[i]+numbers[j] != target {
				j--
			}
		} else if numbers[i]+numbers[j] < target {
			i = sort.SearchInts(numbers[:j], target-numbers[j])
		} else {
			break
		}
	}
	if i == j {
		j++
	}
	return []int{i + 1, j + 1}
}
