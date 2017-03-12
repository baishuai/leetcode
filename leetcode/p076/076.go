package p076

import "fmt"

/**
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

For example,
S = "ADOBECODEBANC"
T = "ABC"
Minimum window is "BANC".

Note:
If there is no such window in S that covers all characters in T, return the empty string "".

If there are multiple such windows, you are guaranteed that there will always be only one unique minimum window in S.
 */

func minWindow(s string, t string) string {
	tHits := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tHits[t[i]] ++
	}
	l, h, hitCounts := 0, 0, 0
	ml, mh, mLen := 0, 0, len(s)+1
	for h < len(s) || hitCounts >= len(t) {

		if hitCounts < len(t) {
			for h < len(s) {
				if _, ok := tHits[s[h]]; ok {
					break
				}
				h++
			}
			if h < len(s) {
				if tHits[s[h]] > 0 {
					hitCounts++
				}
				tHits[s[h]]--
			}
			h++
		} else {
			if v, ok := tHits[s[l]]; ok {
				if v >= 0 {
					hitCounts--
				}
				tHits[s[l]]++
				l++
			}
			for l < len(s) {
				if _, ok := tHits[s[l]]; ok {
					break
				}
				l++
			}
		}
		if hitCounts == len(t) && (h-l) < mLen {
			ml, mh = l, h
			mLen = h - l
			fmt.Println("find", l, h, string([]byte(s)[l:h]))
		}
		if mh-ml == len(t) {
			break
		}
	}
	return string([]byte(s)[ml:mh])
}
