package p075

/**
Given an array with n objects colored red, white or blue, sort them so that objects of the same color are adjacent,
 with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

 */

func sortColors(nums []int) {
	i, j := 0, len(nums)-1
	first_n_mean := len(nums)
	for i < j && i <= first_n_mean {
		for i < j && nums[i] < 1 {
			i++
		}
		for i < j && nums[j] > 1 {
			j--
		}
		if nums[i] > 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j--
			continue
		}
		if nums[j] < 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			continue
		}
		if first_n_mean == len(nums) {
			first_n_mean = j - 1
		}
		for i <= first_n_mean && nums[first_n_mean] == 1 {
			first_n_mean--
		}
		//fmt.Println(i, j, first_n_mean, nums)
		if first_n_mean > i && nums[first_n_mean] < 1 {
			nums[i], nums[first_n_mean] = nums[first_n_mean], nums[i]
			first_n_mean--
			i++
		}
		if first_n_mean > i && nums[first_n_mean] > 1 {
			nums[j], nums[first_n_mean] = nums[first_n_mean], nums[j]
			first_n_mean--
			j--
		}
	}
}
