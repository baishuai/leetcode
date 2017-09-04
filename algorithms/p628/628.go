package p628

/**
Given an integer array, find three numbers whose product is maximum and output the maximum product.

Example 1:
Input: [1,2,3]
Output: 6
Example 2:
Input: [1,2,3,4]
Output: 24
Note:
The length of the given array will be in range [3,104] and all elements are in the range [-1000, 1000].
Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.

*/

type option struct {
	val  int
	none bool
}

func maximumProduct(nums []int) int {

	pop1, pop2, pop3 := option{none: true}, option{none: true}, option{none: true}
	nop1, nop2, nop3 := option{none: true}, option{none: true}, option{none: true}

	for _, n := range nums {
		if n >= 0 {
			if pop1.none || n > pop1.val {
				pop3 = pop2
				pop2 = pop1
				pop1.val = n
				pop1.none = false
			} else if pop2.none || n > pop2.val {
				pop3 = pop2
				pop2.val = n
				pop2.none = false
			} else if pop3.none || n > pop3.val {
				pop3.val = n
				pop3.none = false
			}
		} else {
			if nop1.none || n < nop1.val {
				nop3 = nop2
				nop2 = nop1
				nop1.val = n
				nop1.none = false
			} else if nop2.none || n < nop2.val {
				nop3 = nop2
				nop2.val = n
				nop2.none = false
			} else if nop3.none || n < nop3.val {
				nop3.val = n
				nop3.none = false
			}
		}
	}

	res1 := option{none: true}
	if !pop1.none && !pop2.none && !pop3.none {
		res1.val = pop1.val * pop2.val * pop3.val
		res1.none = false
	}
	res2 := option{none: true}
	if !nop1.none && !nop2.none {
		if !pop1.none {
			res2.none = false
			res2.val = pop1.val * nop1.val * nop2.val
		} else if !nop3.none {
			res2.none = false
			res2.val = nop1.val * nop2.val * nop3.val
		}
	}
	if res1.none {
		return res2.val
	} else if res2.none {
		return res1.val
	}
	if res1.val > res2.val {
		return res1.val
	} else {
		return res2.val
	}
}
