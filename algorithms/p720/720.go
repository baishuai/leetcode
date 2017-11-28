package p720

import (
	"sort"
)

type strs []string

func (s strs) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s strs) Less(i, j int) bool {
	return len(s[i]) < len(s[j]) || (len(s[i]) == len(s[j]) && s[i] < s[j])
}

// Swap swaps the elements with indexes i and j.
func (s strs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func longestWord(words []string) string {
	sort.Sort(strs(words))
	m := make(map[string]bool)
	var ans string
	for _, w := range words {
		if len(w) == 1 {
			m[w] = true
		} else {
			v, ok := m[w[:len(w)-1]]
			m[w] = ok && v
		}
		if m[w] && len(w) > len(ans) {
			ans = w
		}
	}
	return ans
}
