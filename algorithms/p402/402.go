package p402

import (
	"bytes"
)

/**
Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.

Note:
The length of num is less than 10002 and will be ≥ k.
The given num does not contain any leading zero.
Example 1:

Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
Example 2:

Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
Example 3:

Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.
*/

func removeKdigits(num string, k int) string {
	invalid := byte('0') - 1
	numbs := []byte(num)
	for i := 0; i < len(num)-1; i++ {
		if numbs[i] > numbs[i+1] && k > 0 {
			numbs[i] = invalid
			k--
			for j := i - 1; j >= 0 && k > 0; j-- {
				if numbs[j] == invalid {
					continue
				} else if numbs[j] > numbs[i+1] {
					numbs[j] = invalid
					k--
				} else {
					break
				}
			}
		}
	}

	for i := len(numbs) - 1; k > 0 && i >= 0; i-- {
		if numbs[i] != invalid {
			numbs[i] = invalid
			k--
		}
	}
	resBuf := new(bytes.Buffer)
	resBuf.WriteByte('0')
	for i := 0; i < len(numbs); i++ {
		if numbs[i] != invalid {
			resBuf.WriteByte(numbs[i])
		}
	}
	res := resBuf.Bytes()
	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}
	return string(res)
}
