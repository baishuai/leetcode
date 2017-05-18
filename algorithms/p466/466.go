package p466

import (
	"fmt"
)

/**
466. Count The Repetitions

Define S = [s,n] as the string S which consists of n connected strings s. For example, ["abc", 3] ="abcabcabc".

On the other hand, we define that string s1 can be obtained from string s2 if we can remove some characters from s2 such that it becomes s1. For example, “abc” can be obtained from “abdbec” based on our definition, but it can not be obtained from “acbbe”.

You are given two non-empty strings s1 and s2 (each at most 100 characters long) and two integers 0 ≤ n1 ≤ 106 and 1 ≤ n2 ≤ 106. Now consider the strings S1 and S2, where S1=[s1,n1] and S2=[s2,n2]. Find the maximum integer M such that [S2,M] can be obtained from S1.

Example:

Input:
s1="acb", n1=4
s2="ab", n2=2

Return:
2
**/

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	repeatCount := make([]int, len(s2)+2)
	nextIndex := make(map[int]int)
	j, cnt := 0, 0

	for k := 1; k <= n1; k++ {
		for i := 0; i < len(s1); i++ {
			if s1[i] == s2[j] {
				j++
				if j == len(s2) {
					j = 0
					cnt++
				}
			}
		}
		repeatCount[k] = cnt
		if v, ok := nextIndex[j]; ok {
			fmt.Println(nextIndex)
			prefixCount := repeatCount[v]
			patternCount := (repeatCount[k] - repeatCount[v]) * (n1 - v) / (k - v)
			suffixCount := repeatCount[v+(n1-v)%(k-v)] - repeatCount[v]
			return (prefixCount + patternCount + suffixCount) / n2
		}
		nextIndex[j] = k
		fmt.Println(nextIndex)
	}
	return repeatCount[n1] / n2
}
