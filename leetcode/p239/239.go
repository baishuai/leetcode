package p239

/**
Given an array nums, there is a sliding window of size k which is moving
 from the very left of the array to the very right. You can only see the
 k numbers in the window. Each time the sliding window moves right by one position.

For example,
Given nums = [1,3,-1,-3,5,3,6,7], and k = 3.

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Therefore, return the max sliding window as [3,3,5,5,6,7].
 */

func maxSlidingWindow(nums []int, k int) []int {
	maxAns := make([]int, 0)
	deque := make([]int, 0)

	for i, v := range nums {
		if len(deque) != 0 && deque[0] == i-k {
			deque = deque[1:]
		}
		for len(deque) != 0 && nums[deque[len(deque)-1]] < v {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			maxAns = append(maxAns, nums[deque[0]])
		}
	}
	return maxAns
}
