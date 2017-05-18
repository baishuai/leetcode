package p030

//import "fmt"

/**
You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.

For example, given:
s: "barfoothefoobarman"
words: ["foo", "bar"]

You should return the indices: [0,9].
(order does not matter).
 */
// words 会包括重复项

func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)
	wordsLen := len(words)
	if wordsLen <= 0 {
		return ans
	}
	wordSize := len(words[0])
	wordsMap := make(map[string]int)
	wordsCount := make([]int, wordsLen)
	for i, w := range words {
		v, ok := wordsMap[w]
		if ! ok {
			wordsMap[w] = i
			v = i
		}
		wordsCount[v]++
	}
	bs := []byte(s)
	for i := 0; i < wordSize; i++ {
		findSeq := make([]int, 0)
		findIndex := -1
		hitCount := make([]int, wordsLen)
		findCount := 0
		for j := i; j < len(bs)-wordSize+1; j = j + wordSize {
			if v, ok := wordsMap[string(bs[j:j+wordSize])]; ok {

				if hitCount[v] == wordsCount[v] {
					for findSeq[0] != v {
						findCount --
						hitCount[findSeq[0]]--
						findSeq = findSeq[1:]
						findIndex += wordSize
					}
					findCount--
					findSeq = findSeq[1:]
					findIndex += wordSize
				} else {
					hitCount[v]++
				}

				findCount++
				findSeq = append(findSeq, v)
				if findIndex == -1 {
					findIndex = j
				}

				//fmt.Println("j", j, "v", v, "findseq", findSeq)
				//fmt.Println(hitCount)
			} else { // not a word
				findCount = 0
				findSeq = findSeq[0:0]
				for k := 0; k < wordsLen; k++ {
					hitCount[k] = 0
				}
				findIndex = -1
			}

			if findCount == wordsLen {
				ans = append(ans, findIndex)
				//fmt.Println("append ans", findIndex)
				findCount --
				hitCount[findSeq[0]] --
				findSeq = findSeq[1:]
				findIndex += wordSize
			}
		}
	}
	return ans
}
