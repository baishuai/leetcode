package p005

/**
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example:

Input: "babad"

Output: "bab"

Note: "aba" is also a valid answer.
Example:

Input: "cbbd"

Output: "bb"
 */

func longestPalindrome(s string) string {
	sLen := len(s)
	longCount := make([]int, sLen)
	longestIndex := 0
	findInnerLastChar := func(index int) int {
		lastIndex := index
		for s[lastIndex] == s[index] {
			lastIndex ++
			if lastIndex >= sLen {
				break
			}
		}
		return lastIndex - 1
	}

	for i := 0; i < sLen; i++ {
		lastIndex := findInnerLastChar(i)
		var j int
		for j = 1; i-j >= 0 && lastIndex+j < sLen; j++ {
			if s[i-j] != s[lastIndex+j] {
				break
			}
		}
		longCount[i] = (lastIndex - i + 1) + 2*(j-1)
		if longCount[i] > longCount[longestIndex] {
			longestIndex = i
		}
	}
	lastIndex := findInnerLastChar(longestIndex)
	left := longestIndex - (longCount[longestIndex]-(lastIndex-longestIndex+1) )/2
	return string([]byte(s)[left: left+longCount[longestIndex]])
}
