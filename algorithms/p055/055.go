package p055

/**
Given an array of non-negative integers,
 you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

For example:
A = [2,3,1,1,4], return true.

A = [3,2,1,0,4], return false.
*/

func canJump(nums []int) bool {
	furthest := 0
	lens := len(nums)
	for i := 0; i <= furthest && i < lens; i++ {
		if i+nums[i] > furthest {
			furthest = i + nums[i]
		}
	}
	return furthest >= lens-1
}
