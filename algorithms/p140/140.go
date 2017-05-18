package p140

import (
	"strings"
)

/**
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to construct a sentence where each word is a valid dictionary word. You may assume the dictionary does not contain duplicate words.

Return all such possible sentences.

For example, given
s = "catsanddog",
dict = ["cat", "cats", "and", "sand", "dog"].

A solution is ["cats and dog", "cat sand dog"].
 */

type Trie struct {
	root *TrieNode
	res  []string
}

type TrieNode struct {
	word     string
	children []*TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := Trie{root: &TrieNode{children: make([]*TrieNode, 26)}, res: make([]string, 0)}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.root
	for _, v := range []byte(word) {
		if cur.children[v-'a'] == nil {
			cur.children[v-'a'] = &TrieNode{children: make([]*TrieNode, 26)}
		}
		cur = cur.children[v-'a']
	}
	cur.word = word
}

func (this *Trie) find(key []byte, ix int, ws []string, mem []int) int {
	if ix == len(key) {
		this.res = append(this.res, strings.Join(ws, " "))
		return 1
	} else if mem[ix] < 0 {
		return 0
	}
	mem[ix] = 0
	cur := this.root
	for i := ix; i <= len(key); i++ {
		if cur == nil {
			if mem[ix] == 0 {
				mem[ix] = -1
			}
			break
		}
		if cur.word != "" {
			ws = append(ws, cur.word)
			mem[ix] += this.find(key, i, ws, mem)
			ws = ws[:len(ws)-1]
		}
		if i != len(key) {
			cur = cur.children[key[i]-'a']
		}
	}
	if mem[ix] == -1 {
		return 0
	}
	return mem[ix]
}

func wordBreak(s string, wordDict []string) []string {

	trie := Constructor()
	for _, v := range wordDict {
		trie.Insert(v)
	}
	mem := make([]int, len(s))
	trie.find([]byte(s), 0, []string{}, mem)
	return trie.res
}
