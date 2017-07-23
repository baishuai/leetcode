package p648

import "strings"

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	isWord   bool
	word     string
	children []*TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := Trie{root: &TrieNode{children: make([]*TrieNode, 26)}}
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
	cur.isWord = true
	cur.word = word
}

func (this *Trie) Find(key string) string {
	cur := this.root
	for _, v := range []byte(key) {
		if cur == nil {
			break
		}
		if cur.isWord {
			return cur.word
		}
		cur = cur.children[v-'a']
	}
	return key
}

func replaceWords(dict []string, sentence string) string {
	tr := Constructor()
	for _, d := range dict {
		tr.Insert(d)
	}

	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		words[i] = tr.Find(words[i])
	}
	return strings.Join(words, " ")
}
