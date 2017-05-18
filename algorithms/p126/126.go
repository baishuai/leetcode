package p126

/**
Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s) from beginWord to endWord, such that:

Only one letter can be changed at a time
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
For example,

Given:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log","cog"]
Return
  [
    ["hit","hot","dot","dog","cog"],
    ["hit","hot","lot","log","cog"]
  ]
Note:
Return an empty list if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.

 */

type ladderNode struct {
	word string
	used bool
	prev []*ladderNode
}

func (l *ladderNode) addPrev(p *ladderNode) {
	l.prev = append(l.prev, p)
}

func wordsDiff1(a, b string) bool {
	res := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			res++
			if res >= 2 {
				break
			}
		}
	}
	return res == 1
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	//BFS
	bfsQueue := []*ladderNode{{word: beginWord}}

	ladderList := make([]*ladderNode, len(wordList))
	for i, w := range wordList {
		ladderList[i] = &ladderNode{word: w, prev: make([]*ladderNode, 0)}
	}

	find := false
	for len(bfsQueue) > 0 {
		if find {
			break
		}
		preBfsQueue := bfsQueue
		bfsQueue = make([]*ladderNode, 0)

		for _, pnode := range preBfsQueue {
			for _, word := range ladderList {
				if word.used {
					continue
				}
				if wordsDiff1(pnode.word, word.word) {
					if !find && word.word == endWord {
						bfsQueue = bfsQueue[:0]
						find = true
					}
					if !find || (find && word.word == endWord) {
						word.addPrev(pnode)
						if len(word.prev) == 1 {
							bfsQueue = append(bfsQueue, word)
						}
					}
				}
			}
		}
		for _, pnode := range bfsQueue {
			pnode.used = true
		}
	}

	var dfs func(l *ladderNode) [][]string

	dfs = func(l *ladderNode) [][]string {
		if l.prev == nil || len(l.prev) == 0 {
			return [][]string{{l.word}}
		}
		result := make([][]string, 0)
		for _, p := range l.prev {
			pres := dfs(p)
			for _, one := range pres {
				result = append(result, append(one, l.word))
			}
		}
		return result

	}
	res := make([][]string, 0)
	for _, pnode := range bfsQueue {
		res = append(res, dfs(pnode)...)
	}
	return res
}
