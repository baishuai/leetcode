package p067

import (
	"bytes"
)

/**
Given two binary strings, return their sum (also a binary string).

For example,
a = "11"
b = "1"
Return "100".
*/

func addBinary(a string, b string) string {
	carry := byte('0')
	la := len(a) - 1
	lb := len(b) - 1
	ans := bytes.Buffer{}
	for la >= 0 || lb >= 0 || carry > '0' {
		ca := byte('0')
		if la >= 0 {
			ca = a[la]
			la--
		}
		cb := byte('0')
		if lb >= 0 {
			cb = b[lb]
			lb--
		}
		s := (ca + cb + carry) - (3 * '0')
		if s == 1 || s == 0 {
			ans.WriteByte('0' + s)
			carry = '0'
		} else if s == 2 || s == 3 {
			ans.WriteByte('0' + s - 2)
			carry = '1'
		}
	}
	rans := ans.Bytes()
	i, j := 0, len(rans)-1
	for i < j {
		rans[i], rans[j] = rans[j], rans[i]
		i++
		j--
	}
	return string(rans)
}
