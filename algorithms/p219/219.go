package p219

/**
Given an array of integers and an integer k, find out whether there are two distinct indices i and j
 in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.
 */

func containsNearbyDuplicate(nums []int, k int) bool {
	hm := make(map[int]bool, k+1)

	i, j := 0, 0
	for j <= k && j < len(nums) {
		if _, ok := hm[nums[j]]; ok {
			return true
		}
		hm[nums[j]] = true
		j++
	}
	for j < len(nums) {
		delete(hm, nums[i])
		if _, ok := hm[nums[j]]; ok {
			return true
		}
		hm[nums[j]] = true
		i++
		j++
	}
	return false
}
