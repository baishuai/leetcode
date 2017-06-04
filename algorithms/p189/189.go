package p189

/**
Rotate an array of n elements to the right by k steps.

For example, with n = 7 and k = 3,
 the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].
*/

func rotate(nums []int, k int) {

	length := len(nums)
	k = k % length
	i, j := length-k, length-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	i, j = 0, length-k-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	i, j = 0, length-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
