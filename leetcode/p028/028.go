package p028

/**
Implement strStr().

Returns the index of the first occurrence of needle in haystack,
 or -1 if needle is not part of haystack.
 */

func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	for i := 0; i <= m-n; i++ {
		var j int
		for j < n {
			if haystack[i+j] != needle[j] {
				break
			}
			j++
		}
		if j == n {
			return i
		}
	}
	return -1
}
