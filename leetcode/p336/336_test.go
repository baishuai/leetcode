package p336

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	words := []string{"bat", "tab", "cat"}
	res := palindromePairs(words)
	assert.Equal(t, [][]int{{0, 1}, {1, 0}}, res)
}

func Test1(t *testing.T) {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	res := palindromePairs(words)
	assert.Equal(t, [][]int{{0, 1}, {1, 0}, {2, 4}, {3, 2}}, res)
}
