package p039

import (
	"sort"
)

/**
Given a set of candidate numbers (C) (without duplicates) and a target number (T),
 find all unique combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [2, 3, 6, 7] and target 7,
A solution set is:
[
  [7],
  [2, 2, 3]
]
*/

func combinationSum(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	ans := make([][]int, 0)

	one := make([]int, target)
	var comSum func(c, l, t int)
	comSum = func(c, l, t int) {
		if t == 0 {
			onek := make([]int, l)
			copy(onek, one[:l])
			ans = append(ans, onek)
			return
		} else if t < 0 {
			return
		}
		for i, v := range candidates[c:] {
			one[l] = v
			comSum(c+i, l+1, t-v)
		}
	}

	comSum(0, 0, target)
	return ans
}
