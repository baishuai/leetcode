package p164

/**
Given an unsorted array, find the maximum difference between the successive elements in its sorted form.
Try to solve it in linear time/space.
Return 0 if the array contains less than 2 elements.
You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
*/

func maximumGap(nums []int) int {
	//radix sort
	clone := make([]int, len(nums))
	radix := make([]int, len(nums))
	copy(radix, nums)
	length := len(clone)
	last0 := length - 1
	mask := 1
	for mask > 0 && mask < (1<<31)+1 {
		clone, radix = radix, clone
		i, j := 0, length-1

		for ix := 0; ix <= last0; ix++ {
			if (clone[ix] & mask) == 0 {
				radix[i] = clone[ix]
				i++
			} else {
				radix[j] = clone[ix]
				j--
			}
		}
		for ix := length - 1; ix > last0; ix-- {
			if (clone[ix] & mask) == 0 {
				radix[i] = clone[ix]
				i++
			} else {
				radix[j] = clone[ix]
				j--
			}
		}
		last0 = j
		mask <<= 1
	}

	maxGap := 0
	for i := 0; i < length-1; i++ {
		if radix[i+1]-radix[i] > maxGap {
			maxGap = radix[i+1] - radix[i]
		}
	}
	return maxGap
}
