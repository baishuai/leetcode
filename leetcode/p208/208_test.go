package p208

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	trie := Constructor()
	assert.Equal(t, false, trie.StartsWith("go"))
}
