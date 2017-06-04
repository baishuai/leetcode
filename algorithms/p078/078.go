package p078

import (
	"sort"
)

/**
Given a set of distinct integers, nums, return all possible subsets.

Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,3], a solution is:

[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/
// same with subsetsWithDup 090.go
func subsets(nums []int) [][]int {

	sort.Ints(nums)
	ans := make([][]int, 0)
	one := make([]int, len(nums))
	copy(one, nums)
	max := nums[len(nums)-1]
	for len(one) > 0 {
		onecase := make([]int, len(one))
		copy(onecase, one)
		ans = append(ans, onecase)
		v := one[len(one)-1]
		one = one[:len(one)-1]
		if v < max {
			index := sort.SearchInts(nums, v+1)
			for index < len(nums) {
				one = append(one, nums[index])
				index++
			}
		}
	}
	ans = append(ans, one[:])
	return ans
}
