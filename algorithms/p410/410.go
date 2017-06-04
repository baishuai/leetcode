package p410

import (
	"sort"
)

/**
Given an array which consists of non-negative integers and an integer m,
you can split the array into m non-empty continuous subarrays.
Write an algorithm to minimize the largest sum among these m subarrays.

Note:
If n is the length of array, assume the following constraints are satisfied:

1 ≤ n ≤ 1000
1 ≤ m ≤ min(50, n)
Examples:

Input:
nums = [7,2,5,10,8]
m = 2

Output:
18

Explanation:
There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8],
where the largest sum among the two subarrays is only 18.
*/

// returns subarrays count
func splitArrayMax(nums []int64, sum int64) int {
	//nums is sorted
	cntSum := sum
	cnt, idx, pidx := 1, 0, 0
	for idx < len(nums) {
		pidx = idx

		//fmt.Println("search short", nums, cntSum, sort.Search(len(nums[idx:]), func(i int) bool { return nums[i] >= cntSum }))
		idx = sort.Search(len(nums), func(i int) bool { return nums[i] > cntSum })

		if idx < len(nums) && nums[idx] == cntSum {
			idx++
		}
		if pidx == idx {
			cnt = len(nums) + 1
			break
		}
		if idx < len(nums) {
			cnt++
			cntSum = nums[idx-1] + sum
		}
	}
	return cnt
}

func splitArray(nums []int, m int) int {
	if len(nums) == 0 {
		return 0
	}

	nums64 := make([]int64, len(nums))
	nums64[0] = int64(nums[0])
	for i := 1; i < len(nums); i++ {
		nums64[i] = int64(nums64[i-1]) + int64(nums[i])
	}

	minSum := nums64[len(nums64)-1] / int64(m)
	l, r := minSum, nums64[len(nums64)-1]
	for l < r {
		mean := (l + r) >> 1
		if splitArrayMax(nums64, mean) > m {
			l = mean + 1
		} else {
			r = mean
		}

	}
	return int(l)
}
