package p034

/**
Given an array of integers sorted in ascending order,
 find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

For example,
Given [5, 7, 7, 8, 8, 10] and target value 8,
return [3, 4].
 */

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	if len(nums) == 0 {
		return ans
	}
	firstHit := -1
	lo, hi, mid := 0, len(nums)-1, 0
	for lo <= hi {
		mid = (lo + hi ) >> 1
		if nums[mid] == target {
			firstHit = mid
			break
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	if firstHit == -1 {
		return ans
	}

	searchBoundary := func(l, h, t int) int {
		mid := 0
		for l < h {
			mid = (l + h) >> 1
			if nums[mid] < t {
				l = mid + 1
			} else {
				h = mid
			}
		}
		return l
	}

	ans[0] = searchBoundary(lo, firstHit, target)
	ans[1] = searchBoundary(firstHit, hi+1, target+1) - 1
	return ans
}
