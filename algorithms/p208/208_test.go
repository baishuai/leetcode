package p208

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	trie := Constructor()
	assert.Equal(t, false, trie.StartsWith("go"))
}
