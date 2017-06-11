package p414

/**
Given a non-empty array of integers, return the third maximum number in this array.
If it does not exist, return the maximum number. The time complexity must be in O(n).

Example 1:
Input: [3, 2, 1]

Output: 1

Explanation: The third maximum is 1.
Example 2:
Input: [1, 2]

Output: 2

Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
Example 3:
Input: [2, 2, 3, 1]

Output: 1

Explanation: Note that the third maximum here means the third maximum distinct number.
Both numbers with value 2 are both considered as second maximum.
*/

type option struct {
	val  int
	none bool
}

func thirdMax(nums []int) int {
	max := make([]option, 3)
	for i := 0; i < 3; i++ {
		max[i].none = true
	}
outer:
	for _, n := range nums {
		for _, m := range max {
			if !m.none && m.val == n {
				continue outer
			}
		}
		if max[0].none || n > max[0].val {
			max[2] = max[1]
			max[1] = max[0]
			max[0].val = n
			max[0].none = false
		} else if max[1].none || n > max[1].val {
			max[2] = max[1]
			max[1].val = n
			max[1].none = false
		} else if max[2].none || n > max[2].val {
			max[2].val = n
			max[2].none = false
		}
	}

	if max[2].none {
		return max[0].val
	}
	return max[2].val
}
