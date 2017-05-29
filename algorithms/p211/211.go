package p211

/**
Design a data structure that supports the following two operations:

void addWord(word)
bool search(word)
search(word) can search a literal word or a regular expression string containing only letters a-z or .. A . means it can represent any one letter.

For example:

addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
Note:
You may assume that all words are consist of lowercase letters a-z.
**/

type TrieNode struct {
	isWord   bool
	children []*TrieNode
}

type WordDictionary struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	trie := WordDictionary{root: &TrieNode{children: make([]*TrieNode, 26)}}
	return trie
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, v := range []byte(word) {
		if cur.children[v-'a'] == nil {
			cur.children[v-'a'] = &TrieNode{children: make([]*TrieNode, 26)}
		}
		cur = cur.children[v-'a']
	}
	cur.isWord = true
}

func (this *TrieNode) search(words []byte) bool {
	if len(words) == 0 {
		return this.isWord
	}
	left := words[1:]
	b := words[0]
	if b != '.' {
		return this.children != nil &&
			this.children[b-'a'] != nil &&
			this.children[b-'a'].search(left)
	}
	if this.children == nil {
		return false
	}
	for i := 0; i < 26; i++ {
		find := this.children[i] != nil && this.children[i].search(left)
		if find {
			return true
		}
	}
	return false
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.root.search([]byte(word))
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
