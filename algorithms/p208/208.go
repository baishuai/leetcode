package p208

/**
Implement a trie with insert, search, and startsWith methods.
You may assume that all inputs are consist of lowercase letters a-z.
*/

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	isWord   bool
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
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.find(word)
	return cur != nil && cur.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix) != nil
}

func (this *Trie) find(key string) *TrieNode {
	cur := this.root
	for _, v := range []byte(key) {
		if cur == nil {
			break
		}
		cur = cur.children[v-'a']
	}
	return cur
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
