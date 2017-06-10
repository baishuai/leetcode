package p438

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, []int{0, 6}, findAnagrams("cbaebabacd", "abc"))

	assert.Equal(t, []int{0, 1, 2}, findAnagrams("abab", "ab"))
}

func Test1(t *testing.T) {
	assert.Equal(t, []int{}, findAnagrams("", "a"))
}
