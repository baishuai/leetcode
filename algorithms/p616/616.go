package p616

import (
	"bytes"
	"fmt"
)

/**
Given a string s and a list of strings dict, you need to add a closed pair of bold tag <b> and </b> to wrap the substrings in s that exist in dict. If two such substrings overlap, you need to wrap them together by only one pair of closed bold tag. Also, if two substrings wrapped by bold tags are consecutive, you need to combine them.

Example 1:
Input:
s = "abcxyz123"
dict = ["abc","123"]
Output:
"<b>abc</b>xyz<b>123</b>"
Example 2:
Input:
s = "aaabbcc"
dict = ["aaa","aab","bc"]
Output:
"<b>aaabbc</b>c"
*/

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	isWord   bool
	word     string
	children map[byte]*TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := Trie{root: &TrieNode{children: make(map[byte]*TrieNode)}}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.root
	for _, v := range []byte(word) {
		if cur.children[v] == nil {
			cur.children[v] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		cur = cur.children[v]
	}
	cur.isWord = true
	cur.word = word
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix []byte) int {
	node := this.find(prefix)
	if node == nil {
		//fmt.Println("find", prefix, 0)
		return 0
	}
	return len(node.word)
}

func (this *Trie) find(key []byte) *TrieNode {
	cur := this.root

	var f *TrieNode
	for _, v := range key {
		cur = cur.children[v]
		if cur == nil {
			break
		}
		if cur.isWord {
			f = cur
		}
	}
	return f
}

func addBoldTag(s string, dict []string) string {

	wordTrie := Constructor()
	for _, w := range dict {
		wordTrie.Insert(w)
	}
	ss := []byte(s)
	buf := new(bytes.Buffer)
	for i := 0; i < len(s); i++ {

		end := i
		end += wordTrie.StartsWith(ss[i:])
		if end > i {
			buf.WriteString("<b>")
		}
		for end > i {
			c := i + 1
			for c < len(s) && c < end {
				newEnd := c + wordTrie.StartsWith(ss[c:])
				if newEnd > end {
					end = newEnd
				}
				c++
			}
			if end >= len(s) || wordTrie.StartsWith(ss[end:]) == 0 {
				buf.WriteString(s[i:end])
				buf.WriteString("</b>")
				i = end - 1
				break
			} else {
				fmt.Println(wordTrie.StartsWith(ss[end:]))
				end += wordTrie.StartsWith(ss[end:])
			}
		}
		if end == i {
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}
