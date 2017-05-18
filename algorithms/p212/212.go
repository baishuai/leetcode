package p212

/**
Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

For example,
Given words = ["oath","pea","eat","rain"] and board =

[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]
Return ["eat","oath"].

 */

type trieNode struct {
	word     string
	children [26]*trieNode
}

type Trie struct {
	root *trieNode
}

func (t *Trie) insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		v := word[i] - 'a'
		if cur.children[v] == nil {
			cur.children[v] = new(trieNode)
		}
		cur = cur.children[v]
	}
	cur.word = word
}

func dfs(board [][]byte, i, j int, node *trieNode) []string {
	c := board[i][j]
	if c == '#' || node.children[c-'a'] == nil {
		return nil
	}

	node = node.children[c-'a']
	res := make([]string, 0)
	if node.word != "" {
		res = append(res, node.word)
		node.word = ""
	}

	board[i][j] = '#'
	if i > 0 {
		res = append(res, dfs(board, i-1, j, node)...)
	}
	if j > 0 {
		res = append(res, dfs(board, i, j-1, node)...)
	}
	if i < len(board)-1 {
		res = append(res, dfs(board, i+1, j, node)...)
	}
	if j < len(board[i])-1 {
		res = append(res, dfs(board, i, j+1, node)...)
	}
	board[i][j] = c

	return res
}

func findWords(board [][]byte, words []string) []string {

	trie := Trie{root: &trieNode{}}
	for _, w := range words {
		trie.insert(w)
	}
	res := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			res = append(res, dfs(board, i, j, trie.root)...)
		}
	}
	return res
}
