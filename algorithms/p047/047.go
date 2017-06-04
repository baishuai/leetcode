package p047

import (
	"sort"
)

/**
Given a collection of distinct numbers, return all possible permutations.

For example,
[1,2,3] have the following permutations:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

// 字典序全排列生成算法，更快的还有邻位对换法
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	lenn := len(nums)
	nextPermute := func() bool {
		onepermute := make([]int, lenn)
		copy(onepermute, nums)
		ans = append(ans, onepermute)

		for i := lenn - 2; i >= 0; i-- {
			if nums[i] < nums[i+1] {
				j := 0
				for j = lenn - 1; j > i; j-- {
					if nums[j] > nums[i] {
						nums[i], nums[j] = nums[j], nums[i]
						break
					}
				}

				i = i + 1
				j = lenn - 1
				for i < j {
					nums[i], nums[j] = nums[j], nums[i]
					i++
					j--
				}
				return true
			}

		}
		return false
	}

	for nextPermute() {
	}
	return ans
}
