package p014

import (
	"bytes"
	"sort"
)

/**
Write a function to find the longest common prefix string amongst an array of strings.
*/

func longestCommonPrefix(strs []string) string {
	lens := len(strs)
	if lens == 0 {
		return ""
	}
	sort.Strings(strs)
	var buffer bytes.Buffer
	ls, rs := strs[0], strs[lens-1]
	llen, rlen := len(ls), len(rs)
	for i := 0; i < llen && i < rlen; i++ {
		if ls[i] == rs[i] {
			buffer.WriteByte(ls[i])
		} else {
			break
		}
	}
	return buffer.String()
}
