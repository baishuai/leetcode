package p127

/**
Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
For example,

Given:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log","cog"]
As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
return its length 5.
*/

func wordsDiff(a, b string) (res int) {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			res++
		}
	}
	return res
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	//BFS
	length := 1
	bfsQueue := []string{beginWord}

outer:
	for len(bfsQueue) > 0 {
		preBfsQueue := bfsQueue
		bfsQueue = make([]string, 0)
		length++

		for _, str := range preBfsQueue {
			for i, word := range wordList {
				if word == "" {
					continue
				}
				if wordsDiff(str, word) == 1 {
					wordList[i] = ""
					bfsQueue = append(bfsQueue, word)
					if word == endWord {
						break outer
					}
				}
			}
		}
		if len(bfsQueue) == 0 {
			length = 0
			break
		}
	}
	return length
}
