package p327

/**
327. Count of Range Sum

Given an integer array nums, return the number of range sums that lie in [lower, upper] inclusive.
Range sum S(i, j) is defined as the sum of the elements in nums between indices i and j (i ≤ j), inclusive.

Note:
A naive algorithm of O(n2) is trivial. You MUST do better than that.

Example:
Given nums = [-2, 5, -1], lower = -2, upper = 2,
Return 3.
The three ranges are : [0, 0], [2, 2], [0, 2] and their respective sums are: -2, -1, 2.
*/

func countRangeSum(nums []int, lower int, upper int) int {
	sums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	helper := make([]int, len(nums)+1)
	return countWhileMergoSort(sums, helper, 0, len(sums), lower, upper)
}

func countWhileMergoSort(sums, helper []int, start, end int, lower, upper int) int {
	if end-start <= 1 {
		return 0
	}
	mid := start + (end-start)/2
	cnt := countWhileMergoSort(sums, helper, start, mid, lower, upper) +
		countWhileMergoSort(sums, helper, mid, end, lower, upper)

	j, k, t := mid, mid, mid
	for i, r := start, 0; i < mid; i++ {
		for k < end && sums[k]-sums[i] < lower {
			k++
		}
		for j < end && sums[j]-sums[i] <= upper {
			j++
		}
		for t < end && sums[t] < sums[i] {
			helper[r] = sums[t]
			r++
			t++
		}
		helper[r] = sums[i]
		r++
		cnt += j - k
	}
	copy(sums[start:], helper[:t-start])
	return cnt
}
