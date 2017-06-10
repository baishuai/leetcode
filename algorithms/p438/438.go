package p438

/**
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
*/

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	sig := make([]int, 26)
	window := make([]int, 26)
	for i := 0; i < len(p); i++ {
		sig[p[i]-'a'] += 1
		window[s[i]-'a'] += 1
	}

	eq := func(a, b []int, l int) bool {
		for c := 0; c < l; c++ {
			if a[c] != b[c] {
				return false
			}
		}
		return true
	}

	res := make([]int, 0)
	for l, r := 0, len(p); r <= len(s); l, r = l+1, r+1 {
		if eq(window, sig, 26) {
			res = append(res, l)
		}
		if r < len(s) {
			window[s[l]-'a'] -= 1
			window[s[r]-'a'] += 1
		}
	}
	return res
}
