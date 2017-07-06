package p567

/**
Given two strings s1 and s2, write a function to return true if s2 contains
the permutation of s1. In other words, one of the first string's permutations
is the substring of the second string.

Example 1:
Input:s1 = "ab" s2 = "eidbaooo"
Output:True
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:
Input:s1= "ab" s2 = "eidboaoo"
Output: False
 */

// two pointer

func checkInclusion(s1 string, s2 string) bool {
	charCnt := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		charCnt[s1[i]-'a'] ++
	}
	cnt := len(s1)

	i, j := 0, 0

	for j < len(s2) {
		c := s2[j] - 'a'
		if charCnt[c] > 0 {
			charCnt[c] --
			cnt--
			if cnt == 0 {
				break
			}
		} else {
			ic := s2[i] - 'a'
			for ic != c {
				charCnt[ic]++
				cnt++
				i++
				ic = s2[i] - 'a'
			}
			i++
		}
		j++
	}
	return cnt == 0
}
