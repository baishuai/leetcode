package p624

import "math"

/**
Given m arrays, and each array is sorted in ascending order.
 Now you can pick up two integers from two different arrays (each array picks one) and calculate the distance.
  We define the distance between two integers a and b to be their absolute difference |a-b|.
  Your task is to find the maximum distance.

Example 1:
Input:
[[1,2,3],
 [4,5],
 [1,2,3]]
Output: 4
Explanation:
One way to reach the maximum distance 4 is to pick 1 in the first or third array and pick 5 in the second array.
Note:
Each given array will have at least 1 number. There will be at least two non-empty arrays.
The total number of the integers in all the m arrays will be in the range of [2, 10000].
The integers in the m arrays will be in the range of [-10000, 10000].
 */

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func maxDistance(arrays [][]int) int {

	if arrays == nil || len(arrays) == 0 || arrays[0] == nil || len(arrays[0]) == 0 {
		return 0
	}

	min := arrays[0][0]
	max := arrays[0][len(arrays[0])-1]

	diff := math.MinInt32

	for i := 1; i < len(arrays); i++ {
		head := arrays[i][0]
		tail := arrays[i][len(arrays[i])-1]
		if Abs(max-head) > diff {
			diff = Abs(max - head)
		}
		if Abs(tail-min) > diff {
			diff = Abs(tail - min)
		}
		max = Max(max, tail)
		min = Min(min, head)
	}
	return diff
}
