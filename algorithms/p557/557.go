package p557

import "bytes"

/**
Given a string, you need to reverse the order of characters in each word within a sentence
 while still preserving whitespace and initial word order.

Example 1:
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
Note: In the string, each word is separated by single space and there will not be any
extra space in the string.
*/

func reverse(b []byte) {
	i, j := 0, len(b)-1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}

func reverseWords(s string) string {

	sbytes := []byte(s)
	pIndex := 0
	index := 0
	for pIndex < len(sbytes) {
		index = bytes.IndexByte(sbytes[pIndex:], ' ')
		if index == -1 {
			index = len(sbytes)
		} else {
			index += pIndex
		}
		reverse(sbytes[pIndex:index])
		pIndex = index + 1
	}
	return string(sbytes)
}
