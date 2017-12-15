package p724

func pivotIndex(nums []int) int {
	s := 0
	for i, v := range nums {
		nums[i] += s
		s += v
	}

	n := len(nums)
	for i := 0; i < n; i++ {
		left := 0
		if i > 0 {
			left = nums[i-1]
		}
		right := nums[n-1] - nums[i]
		if left == right {
			return i
		}
	}
	return -1
}
