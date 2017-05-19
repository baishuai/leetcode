package p214

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, "aaacecaaa", shortestPalindrome("aacecaaa"))
}

func Test1(t *testing.T) {
	assert.Equal(t, "dcbabcd", shortestPalindrome("abcd"))
}

func Test2(t *testing.T) {
	assert.Equal(t, "", shortestPalindrome(""))
}

func Test3(t *testing.T) {
	assert.Equal(t, "bbabb", shortestPalindrome("abb"))
}
