package p003

/**
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

Subscribe to see which companies asked this question.
*/

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	hmap := make(map[byte]int)
	max := 0
	count := 0
	lastIndex := 0
	for i := 0; i < length; i++ {
		if v, ok := hmap[s[i]]; ok {
			//重复，从下一个开始
			if count > max {
				max = count
			}
			for j := lastIndex; j < v; j++ {
				delete(hmap, s[j])
			}
			lastIndex = v + 1
			count = i - lastIndex + 1
			hmap[s[i]] = i
			continue
		}
		hmap[s[i]] = i
		count++
	}
	if count > max {
		max = count
	}
	return max
}
