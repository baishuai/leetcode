package p091

/**
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given an encoded message containing digits, determine the total number of ways to decode it.

For example,
Given encoded message "12", it could be decoded as "AB" (1 2) or "L" (12).

The number of ways decoding "12" is 2.
 */

func numDecodings(s string) int {
	ss := []byte(s)
	if len(ss) == 0 {
		return 0
	} else if len(ss) == 1 {
		if ss[0] != '0' {
			return 1
		} else {
			return 0
		}
	}
	dps := make([]int, len(ss)+1)
	dps[0], dps[1] = 1, 1
	if ss[0] == '0' {
		dps[1] = 0
	}
	for i := 2; i <= len(ss); i++ {
		if ss[i-1] != '0' {
			dps[i] += dps[i-1]
		}
		if ss[i-2] == '1' || (ss[i-2] == '2' && ss[i-1] <= '6') {
			dps[i] += dps[i-2]
		}
	}
	return dps[len(ss)]
}
