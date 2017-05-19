package p214

import (
	"bytes"
)

/**
Given a string S, you are allowed to convert it to a palindrome by adding characters in front of it. Find and return the shortest palindrome you can find by performing this transformation.

For example:

Given "aacecaaa", return "aaacecaaa".

Given "abcd", return "dcbabcd".

*/

// todo  use kmp
func shortestPalindrome(s string) string {
	ss := []byte(s)

	length := len(ss)
	if length <= 1 {
		return s
	}
	ls, rs := 0, 0

	minls, minrs := 0, 0
	minPadding := length + 1
	for rs < length/2 {
		l, r := ls, rs
		for {
			if l < 0 || r >= length {
				if l < 0 && length-r < minPadding {
					minls, minrs = ls, rs
					minPadding = length - r
				} else if r >= length && l+1 < minPadding {
					minls, minrs = ls, rs
					minPadding = l + 1
				}
				break
			}
			if ss[l] == ss[r] {
				l--
				r++
			} else {
				break
			}
		}
		if rs == ls {
			rs++
		} else {
			ls++
		}
	}

	var buf bytes.Buffer
	if (minls - 0 + 1) == (length - minrs) {
		buf.WriteString(s)
	} else if (minls - 0 + 1) > (length - minrs) {
		buf.WriteString(s)
		itx := minls - (length - minrs)
		for itx >= 0 {
			buf.WriteByte(ss[itx])
			itx--
		}
	} else {
		itx := length - 1
		end := minrs + minls
		for itx > end {
			buf.WriteByte(ss[itx])
			itx--
		}
		buf.WriteString(s)
	}
	return buf.String()
}
