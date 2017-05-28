package p600

/**
Given a positive integer n, find the number of non-negative integers less than or equal to n, whose binary representations do NOT contain consecutive ones.

Example 1:
Input: 5
Output: 5
Explanation:
Here are the non-negative integers <= 5 with their corresponding binary representations:
0 : 0
1 : 1
2 : 10
3 : 11
4 : 100
5 : 101
Among them, only integer 3 disobeys the rule (two consecutive ones) and the other 5 satisfy the rule.
Note: 1 <= n <= 109
**/

//dps1(i) dps0(i) 表示1或0开头的 i 位二进制数没有相连1的数的个数
func findIntegers(num int) int {
	nbits := 0
	for n := num; n > 0; n = n >> 1 {
		nbits++
	}
	dps0, dps1 := make([]int, nbits+1), make([]int, nbits+1)
	dps0[1], dps1[1] = 1, 1
	for i := 2; i <= nbits; i++ {
		dps0[i] = dps0[i-1] + dps1[i-1]
		dps1[i] = dps0[i-1]
	}
	result := dps0[nbits] + dps1[nbits]
	pre1 := false
	for i := uint(nbits); i > 0; i-- {
		mask := 1 << (i - 1)
		if mask&num > 0 {
			if pre1 {
				break
			} else {
				pre1 = true
			}
		} else {
			if pre1 {
				pre1 = false
				continue
			} else {
				result -= dps1[i]
			}
		}
	}
	return result
}
