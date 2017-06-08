package p516

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 4, longestPalindromeSubseq("bbbab"))
	assert.Equal(t, 2, longestPalindromeSubseq("cbbd"))
}
