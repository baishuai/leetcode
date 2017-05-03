package p472

/**
Given a list of words (without duplicates), please write a program that returns all concatenated words in the given list of words.

A concatenated word is defined as a string that is comprised entirely of at least two shorter words in the given array.

Example:
Input: ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]

Output: ["catsdogcats","dogcatsdog","ratcatdogcat"]

Explanation: "catsdogcats" can be concatenated by "cats", "dog" and "cats";
 "dogcatsdog" can be concatenated by "dog", "cats" and "dog";
"ratcatdogcat" can be concatenated by "rat", "cat", "dog" and "cat".
Note:
The number of elements of the given array will not exceed 10,000
The length sum of elements in the given array will not exceed 600,000.
All the input string will only include lower case letters.
The returned elements order does not matter.
 */

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	word     bool
	children []*TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := Trie{root: &TrieNode{children: make([]*TrieNode, 26)}}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	cur := this.root
	for _, v := range []byte(word) {
		if cur.children[v-'a'] == nil {
			cur.children[v-'a'] = &TrieNode{children: make([]*TrieNode, 26)}
		}
		cur = cur.children[v-'a']
	}
	cur.word = true
}

func (this *Trie) isConcatenated(key []byte, ix int) bool {
	concat := false
	cur := this.root
	for i := ix; i <= len(key); i++ {
		if cur == nil {
			break
		}
		if i == len(key) {
			concat = concat || (cur.word && ix != 0)
			continue
		}
		if cur.word && this.isConcatenated(key, i) {
			concat = true
			break
		}
		cur = cur.children[key[i]-'a']
	}
	return concat
}

func findAllConcatenatedWordsInADict(words []string) []string {
	trie := Constructor()
	for _, v := range words {
		trie.Insert(v)
	}

	res := make([]string, 0)
	for _, w := range words {
		if trie.isConcatenated([]byte(w), 0) {
			res = append(res, w)
		}
	}
	return res
}
