package p744

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	letters := []byte{'c', 'f', 'j'}
	assert.Equal(t, byte('c'), nextGreatestLetter(letters, 'a'))
}

func Test1(t *testing.T) {
	letters := []byte{'c', 'f', 'j'}
	assert.Equal(t, byte('c'), nextGreatestLetter(letters, 'j'))
}
