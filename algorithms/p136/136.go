package p136

/**
Given an array of integers, every element appears twice except for one.
 Find that single one.

Note:
Your algorithm should have a linear runtime complexity.
 Could you implement it without using extra memory?
*/

func singleNumber(nums []int) int {
	ans := 0
	for _, n := range nums {
		ans ^= n
	}
	return ans
}

func singleNumber_slow(nums []int) int {
	hmap := make(map[int]bool)
	for _, num := range nums {
		hmap[num] = !hmap[num]
	}
	ans := -1
	for k, v := range hmap {
		if v {
			ans = k
			break
		}
	}
	return ans
}
