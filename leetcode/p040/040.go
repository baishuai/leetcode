package p040

import (
	"sort"
)

/**
Given a collection of candidate numbers (C) and a target number (T),
 find all unique combinations in C where the candidate numbers sums to T.

Each number in C may only be used once in the combination.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [10, 1, 2, 7, 6, 1, 5] and target 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
 */

func combinationSum2(candidates []int, target int) [][]int {

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
			if i == 0 || (i > 0 && candidates[c+i-1] != v) {
				comSum(c+i+1, l+1, t-v)
			}
		}
	}
	comSum(0, 0, target)

	return ans
}
