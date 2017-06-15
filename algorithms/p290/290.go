package p290

import "strings"

/**
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter
in pattern and a non-empty word in str.

Examples:
pattern = "abba", str = "dog cat cat dog" should return true.
pattern = "abba", str = "dog cat cat fish" should return false.
pattern = "aaaa", str = "dog cat cat dog" should return false.
pattern = "abba", str = "dog dog dog dog" should return false.
Notes:
You may assume pattern contains only lowercase letters, and str contains lowercase
letters separated by a single space.
*/

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern) {
		return false
	}
	c2s := make(map[byte]string)
	strset := make(map[string]bool)
	for i := 0; i < len(pattern); i++ {
		if v, ok := c2s[pattern[i]]; ok {
			if v != strs[i] {
				return false
			}
		} else if strset[strs[i]] == false {
			c2s[pattern[i]] = strs[i]
			strset[strs[i]] = true
		} else {
			return false
		}
	}
	return true
}
