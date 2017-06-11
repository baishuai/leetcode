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
		if cur.children[v-'a'] == nil {
			cur.children[v-'a'] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		cur = cur.children[v-'a']
	}
	cur.isWord = true
	cur.word = word
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.find(word)
	return cur != nil && cur.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) int {
	node := this.find(prefix)
	if node == nil {
		return 0
		fmt.Println("find", prefix, 0)

	}
	fmt.Println("find", prefix, len(node.word))

	return len(node.word)
}

func (this *Trie) find(key string) *TrieNode {
	cur := this.root

	var f *TrieNode
	for _, v := range []byte(key) {
		if cur == nil {
			break
		}
		if cur.isWord {
			f = cur
		}
		cur = cur.children[v-'a']
	}
	return f
}

func addBoldTag(s string, dict []string) string {

	wordTrie := Constructor()
	for _, w := range dict {
		wordTrie.Insert(w)
	}

	buf := new(bytes.Buffer)
	for i := 0; i < len(s); i++ {

		end := i

		j := i + 100
		if j > len(s) {
			j = len(s)
		}
		end += wordTrie.StartsWith(s[i:j])

		if end > i {
			buf.WriteString("<b>")
			fmt.Println("loop end", end)
		}
		for end > i {
			c := i + 1
			for c < len(s) && c < end {
				j = c + 100
				if j > len(s) {
					j = len(s)
				}
				newEnd := c + wordTrie.StartsWith(s[c:j])
				fmt.Println("newend", newEnd, c)
				if newEnd > end {
					end = newEnd
				}
				c++
				fmt.Println("loop c", end)
			}
			if end >= len(s) || wordTrie.StartsWith(s[end:]) == 0 {
				buf.WriteString(s[i:end])
				buf.WriteString("</b>")
				i = end - 1
				break
			} else {
				fmt.Println(wordTrie.StartsWith(s[end:]))
				end += wordTrie.StartsWith(s[end:])
			}
		}
		if end == i {
			buf.WriteByte(s[i])
		}

	}
	return buf.String()
}
