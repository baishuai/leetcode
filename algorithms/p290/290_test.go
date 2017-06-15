package p290

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, true, wordPattern("aba", "as dfs as"))
	assert.Equal(t, false, wordPattern("as", "sd df dfg"))
	assert.Equal(t, false, wordPattern("aa", "aa bb"))
	assert.Equal(t, false, wordPattern("ab", "aa aa"))
}
