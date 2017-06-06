package p467

/**
Consider the string s to be the infinite wraparound string of "abcdefghijklmnopqrstuvwxyz",
so s will look like this: "...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".

Now we have another string p. Your job is to find out how many unique non-empty substrings of p are present in s.
In particular, your input is the string p and you need to output the number of different non-empty substrings of p in the string s.

Note: p consists of only lowercase English letters and the size of p might be over 10000.

Example 1:
Input: "a"
Output: 1

Explanation: Only the substring "a" of string "a" is in the string s.
Example 2:
Input: "cac"
Output: 2
Explanation: There are two substrings "a", "c" of string "cac" in the string s.
Example 3:
Input: "zab"
Output: 6
Explanation: There are six substrings "z", "a", "b", "za", "ab", "zab" of string "zab" in the string s.
*/

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func findSubstringInWraproundString(p string) int {
	cnt := make([]int, 26)

	curLen := 0
	for i := 0; i < len(p); i++ {
		if i > 0 && (p[i]-p[i-1] == 1 || p[i-1]-p[i] == 25) {
			curLen++
		} else {
			curLen = 1
		}
		index := p[i] - 'a'
		cnt[index] = max(cnt[index], curLen)
	}
	sum := 0
	for _, v := range cnt {
		sum += v
	}
	return sum
}
