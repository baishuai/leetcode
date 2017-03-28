package p154

func findMin(nums []int) int {

	lo, hi, mid := 0, len(nums)-1, 0
	for lo < hi-1 {
		mid = (lo + hi) / 2
		if nums[lo] < nums[mid] {
			lo = mid
		} else if nums[lo] > nums[mid] {
			hi = mid
		} else if nums[lo] == nums[lo+1] {
			lo++
		} else if nums[hi] == nums[hi-1] {
			hi--
		}
	}
	lo++
	min := nums[0]
	if lo < len(nums) && nums[lo] < nums[0] {
		min = nums[lo]
	}
	return min
}
