package p205

/**
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

For example,
Given "egg", "add", return true.

Given "foo", "bar", return false.

Given "paper", "title", return true.
*/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hms2t := make(map[byte]byte)
	hmt2s := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		vs2t, oks2s := hms2t[s[i]]
		vt2s, okt2s := hmt2s[t[i]]
		if !oks2s && !okt2s {
			hms2t[s[i]] = t[i]
			hmt2s[t[i]] = s[i]
		} else if okt2s && oks2s {
			if vs2t != t[i] || vt2s != s[i] {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
