package p080

/**
Follow up for "Remove Duplicates":
What if duplicates are allowed at most twice?

For example,
Given sorted array nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of
 nums being 1, 1, 2, 2 and 3. It doesn't matter what you leave beyond the new length.
 */

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	prev := nums[0]
	length := 1
	index := 1
	dup := false
	for _, v := range nums[1:] {
		if !dup && v == prev {
			nums[index] = v
			index++
			length++
			dup = true
		} else if v != prev {
			prev = v
			nums[index] = v
			index++
			length++
			dup = false
		}
	}
	return length
}
