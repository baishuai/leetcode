package p300

import "sort"

/**
Given an unsorted array of integers, find the length of longest increasing subsequence.

For example,
Given [10, 9, 2, 5, 3, 7, 101, 18],
The longest increasing subsequence is [2, 3, 7, 101], therefore the length is 4.
Note that there may be more than one LIS combination, it is only necessary for you to return the length.

Your algorithm should run in O(n^2) complexity.
*/

func lengthOfLIS(nums []int) int {
	incSeq := make([]int, 0)
	for _, v := range nums {
		idx := sort.SearchInts(incSeq, v)
		if idx < len(incSeq) {
			incSeq[idx] = v
		} else {
			incSeq = append(incSeq, v)
		}
	}
	return len(incSeq)
}
