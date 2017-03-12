package p049

import (
	"sort"
)

/**
Given an array of strings, group anagrams together.

For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
Return:

[
  ["ate", "eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note: All inputs will be in lower-case.
 */


func groupAnagrams(strs []string) [][]string {
	groupCount := 0
	strsHmap := make(map[int]int)
	ans := make([][]string, 0)

	prime := []byte{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

	for _, str := range strs {
		skey := 1
		for _, c := range []byte(str) {
			skey *= int(prime[c-'a'])
		}

		if v, ok := strsHmap[skey]; ok {
			ans[v] = append(ans[v], str)
		} else {
			strsHmap[skey] = groupCount
			ans = append(ans, []string{str})
			groupCount++
		}
	}
	return ans
}
