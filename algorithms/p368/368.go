package p368

import (
	"sort"
)

/**
Given a set of distinct positive integers, find the largest subset such that every pair (Si, Sj) of elements in this subset satisfies: Si % Sj = 0 or Sj % Si = 0.

If there are multiple solutions, return any subset is fine.

Example 1:

nums: [1,2,3]

Result: [1,2] (of course, [1,3] will also be ok)
Example 2:

nums: [1,2,4,8]

Result: [1,2,4,8]
*/

// 复杂度 O(n^2)

func largestDivisibleSubset(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{}
	}
	sort.Ints(nums)

	preNums := make([]int, len(nums))
	divCnt := make([]int, len(nums))
	divCnt[0] = 1
	globalMaxi := 0

	for i := 1; i < len(nums); i++ {
		v := nums[i]
		maxv := 1
		for j := i - 1; j >= 0; j-- {
			if v%nums[j] == 0 && divCnt[j]+1 > maxv {
				maxv = divCnt[j] + 1
				preNums[i] = j
			}
		}
		divCnt[i] = maxv
		if maxv > divCnt[globalMaxi] {
			globalMaxi = i
		}
	}
	res := make([]int, divCnt[globalMaxi])
	cur := globalMaxi
	for i := divCnt[cur] - 1; i >= 0; i-- {
		res[i] = nums[cur]
		cur = preNums[cur]
	}
	return res
}
