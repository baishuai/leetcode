package p336

/**
Given a list of unique words, find all pairs of distinct indices (i, j) in the given list, so that the concatenation of the two words, i.e. words[i] + words[j] is a palindrome.

Example 1:
Given words = ["bat", "tab", "cat"]
Return [[0, 1], [1, 0]]
The palindromes are ["battab", "tabbat"]
Example 2:
Given words = ["abcd", "dcba", "lls", "s", "sssll"]
Return [[0, 1], [1, 0], [3, 2], [2, 4]]
The palindromes are ["dcbaabcd", "abcddcba", "slls", "llssssll"]
*/

type trieNode struct {
	word     int
	children [26]*trieNode
}

type Trie struct {
	root *trieNode
}

func (n *trieNode) dfsFind(s []byte) []int {
	res := make([]int, 0)
	if n.word > 0 && isPalindrome(s) {
		res = append(res, n.word-1)
	}
	for i := byte(0); i < 26; i++ {
		if n.children[i] != nil {
			res = append(res, n.children[i].dfsFind(append(s, 'a'+i))...)
		}
	}
	return res
}

func (t *Trie) insert(word []byte, ix int) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		v := word[i] - 'a'
		if cur.children[v] == nil {
			cur.children[v] = new(trieNode)
		}
		cur = cur.children[v]
	}
	cur.word = ix
}

func (t *Trie) search(word []byte) []int {
	ix := 0
	node := t.root
	res := make([]int, 0)
	for ix < len(word) && node != nil {
		if node.word > 0 && isPalindrome(word[ix:]) {
			res = append(res, node.word-1)
		}
		if next := node.children[word[ix]-'a']; next != nil {
			ix++
			node = next
		} else {
			node = nil
			break
		}
	}
	if ix == len(word) && node != nil {
		res = append(res, node.dfsFind([]byte{})...)
	}
	return res
}

func isPalindrome(w []byte) bool {
	i, j := 0, len(w)-1
	for i <= j {
		if w[i] == w[j] {
			i++
			j--
		} else {
			break
		}
	}
	return i > j
}

func reverse(s string) []byte {
	res := make([]byte, len(s))
	i, j := 0, len(s)-1
	for i <= j {
		res[i] = s[j]
		res[j] = s[i]
		i++
		j--
	}
	return res
}

func palindromePairs(words []string) [][]int {
	trie := Trie{root: &trieNode{}}

	for i, v := range words {
		trie.insert(reverse(v), i+1)
	}

	res := make([][]int, 0)
	for i, v := range words {
		pairs := trie.search([]byte(v))
		for _, p := range pairs {
			if p != i {
				res = append(res, []int{i, p})
			}
		}
	}

	return res
}
