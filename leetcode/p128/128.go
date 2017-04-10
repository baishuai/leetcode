package p128

/**
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

For example,
Given [100, 4, 200, 1, 3, 2],
The longest consecutive elements sequence is [1, 2, 3, 4]. Return its length: 4.

Your algorithm should run in O(n) complexity.
 */

func longestConsecutive(nums []int) int {
	res := 0
	hm := make(map[int]int)
	for _, v := range nums {

		if _, ok := hm[v]; !ok {
			left := hm[v-1]
			right := hm[v+1]
			sum := left + right + 1
			hm[v] = sum
			if sum > res {
				res = sum
			}

			hm[v-left] = sum
			hm[v+right] = sum
		}
	}
	return res
}
