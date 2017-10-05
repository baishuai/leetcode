package p434

import "strings"

func countSegments(s string) int {
	ss := strings.Split(s, " ")
	cnt := 0
	for _, sp := range ss {
		if len(sp) > 0 {
			cnt++
		}
	}
	return cnt;
}
