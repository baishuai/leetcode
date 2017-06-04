package p209

/**
Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

For example, given the array [2,3,1,2,4,3] and s = 7,
the subarray [4,3] has the minimal length under the problem constraint.
*/

func minSubArrayLen(s int, nums []int) int {
	i, j := 0, 0
	length := len(nums)
	if length == 0 {
		return 0
	}
	partSum := nums[0]
	minLen := length + 1
	for {
		if partSum >= s {
			if j-i+1 < minLen {
				minLen = j - i + 1
			}
		}
		if (j == length-1 && partSum < s) || (i == length) {
			break
		}
		if partSum < s {
			j++
			partSum += nums[j]
		} else {
			partSum -= nums[i]
			i++
		}
	}
	return minLen % (length + 1)
}
