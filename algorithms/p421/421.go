package p421

type TrieNode struct {
	next [2]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	trie := Trie{root: new(TrieNode)}
	return trie
}

func (t *Trie) Insert(num int) {
	cur := t.root
	for i := 31; i >= 0; i-- {
		index := (num >> uint(i) & 1)
		if cur.next[index] == nil {
			cur.next[index] = new(TrieNode)
		}
		cur = cur.next[index]
	}
}

func (t *Trie) MaxXorWith(num int) int {
	res := 0
	cur := t.root
	for i := 31; i >= 0; i-- {
		index := 1 - (num >> uint(i) & 1)
		if cur.next[index] != nil {
			res <<= 1
			res++
			cur = cur.next[index]
		} else {
			res <<= 1
			cur = cur.next[1-index]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		b = a
	}
	return b
}

func findMaximumXOR(nums []int) int {

	t := Constructor()

	for _, n := range nums {
		t.Insert(n)
	}

	res := 0
	for _, n := range nums {
		res = max(res, t.MaxXorWith(n))
	}
	return res
}
