package p523

/**
Given a list of non-negative numbers and a target integer k, write a function to check if the array has a continuous subarray of size at least 2 that sums up to the multiple of k, that is, sums up to n*k where n is also an integer.

Example 1:
Input: [23, 2, 4, 6, 7],  k=6
Output: True
Explanation: Because [2, 4] is a continuous subarray of size 2 and sums up to 6.
Example 2:
Input: [23, 2, 6, 4, 7],  k=6
Output: True
Explanation: Because [23, 2, 6, 4, 7] is an continuous subarray of size 5 and sums up to 42.
**/

func checkSubarraySum(nums []int, k int) bool {
	mods := make(map[int]int)
	sum := 0
	mods[sum] = -1
	for i, n := range nums {
		sum = (sum + n)
		if k != 0 {
			sum %= k
		}
		if v, ok := mods[sum]; ok {
			if i-v >= 2 {
				return true
			}
			continue
		}
		mods[sum] = i
	}
	return false
}
