package p228

import "fmt"

/**
Given a sorted integer array without duplicates, return the summary of its ranges.

For example, given [0,1,2,4,5,7], return ["0->2","4->5","7"].
*/

func summaryRanges(nums []int) []string {

	res := make([]string, 0)
	if nums == nil || len(nums) == 0 {
		return res
	}

	l := nums[0]
	r := l - 1
	for i := 1; i <= len(nums); i++ {
		if i < len(nums) && nums[i]-nums[i-1] == 1 {
			r = nums[i]
		} else {
			if r > l {
				res = append(res, fmt.Sprintf("%d->%d", l, r))
			} else {
				res = append(res, fmt.Sprintf("%d", l))
			}
			if i < len(nums) {
				l = nums[i]
				r = l - 1
			}
		}
	}
	return res
}
