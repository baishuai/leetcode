package p720

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	assert.Equal(t, "apple", longestWord(words))
}
