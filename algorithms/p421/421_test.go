package p421

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	trie := new(TrieNode)
	assert.Equal(t, 2, len(trie.next))
}

func Test1(t *testing.T) {

	nums := []int{3, 10, 5, 25, 2, 8}
	assert.Equal(t, 28, findMaximumXOR(nums))
}
