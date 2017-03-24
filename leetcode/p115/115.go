package p115

import "fmt"

/**
Given a string S and a string T, count the number of distinct subsequences of T in S.

A subsequence of a string is a new string which is formed from the original string
 by deleting some (can be none) of the characters without disturbing the relative positions
  of the remaining characters. (ie, "ACE" is a subsequence of "ABCDE" while "AEC" is not).

Here is an example:
S = "rabbbit", T = "rabbit"

Return 3.
 */

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	if len(t) == 0 {
		return 1
	}

	dpPre := make([]int, len(s))
	dpNow := make([]int, len(s))
	for j := 0; j < len(s); j++ {
		dpNow[j] = 1
	}
	for i := 0; i < len(t); i++ {
		dpNow, dpPre = dpPre, dpNow
		tc := t[i]
		if s[0] == tc && i == 0 {
			dpNow[0] = 1
		} else {
			dpNow[0] = 0
		}
		for j := 1; j < len(s); j++ {
			if s[j] == tc {
				dpNow[j] = dpNow[j-1] + dpPre[j-1]
			} else {
				dpNow[j] = dpNow[j-1]
			}
		}
		fmt.Println(dpNow)

	}
	return dpNow[len(s)-1]
}
