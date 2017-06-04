package p045

/**
Given an array of non-negative integers, you are initially
 positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

For example:
Given array A = [2,3,1,1,4]

The minimum number of jumps to reach the last index is 2.
 (Jump 1 step from index 0 to 1, then 3 steps to the last index.)
*/

func jump(nums []int) int {
	furthest, step := 0, 0
	lens := len(nums)
	i, j := 0, 0
	for furthest < lens-1 {
		j = furthest
		for i <= j {
			if i+nums[i] > furthest {
				furthest = i + nums[i]
			}
			i++
		}
		step++
	}
	return step
}
