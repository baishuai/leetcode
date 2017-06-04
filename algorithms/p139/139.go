package p139

/**
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
determine if s can be segmented into a space-separated sequence of one or more dictionary words.
You may assume the dictionary does not contain duplicate words.

For example, given
s = "leetcode",
dict = ["leet", "code"].

Return true because "leetcode" can be segmented as "leet code".
*/

func wordBreak(s string, wordDict []string) bool {
	len2rowMap := make(map[int]int)
	wordsByLen := make([][]string, 0)
	baseLen, maxLen := 0, 0
	for _, v := range wordDict {
		if len(v) > len(s) {
			continue
		}
		i, ok := len2rowMap[len(v)]
		if !ok {
			wordsByLen = append(wordsByLen, make([]string, 0))
			i = baseLen
			len2rowMap[len(v)] = i
			baseLen++
			if len(v) > maxLen {
				maxLen = len(v)
			}
		}
		wordsByLen[i] = append(wordsByLen[i], v)
	}

	bs := []byte(s)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	base := 0
	for base = 0; base <= len(s); base++ {
		if !dp[base] {
			continue
		}
		for l := 1; l <= maxLen; l++ {
			if i, ok := len2rowMap[l]; ok {
				pred := base + l
				if pred > len(s) || dp[pred] {
					continue
				}
				words := wordsByLen[i]
				partS := string(bs[base:pred])
				for _, w := range words {
					if w == partS {
						dp[pred] = true
						break
					}
				}
			}
		}
	}
	return dp[len(s)]
}
