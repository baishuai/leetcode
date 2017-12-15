package p719

import (
	"sort"
)

func countDistance(nums []int, d int) int {
	//slide window
	l, r := 0, 0
	cnt := 0
	for r < len(nums) {
		for r < len(nums) && nums[r]-nums[l] <= d {
			cnt += r - l
			r++
		}
		if r < len(nums) {
			for nums[r]-nums[l] > d {
				l++
			}
		}
	}
	return cnt
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	if len(nums) <= 1 {
		return 0
	}
	n := nums[len(nums)-1] - nums[0]
	v := sort.Search(n, func(i int) bool {
		return countDistance(nums, i) >= k
	})
	return v
}
