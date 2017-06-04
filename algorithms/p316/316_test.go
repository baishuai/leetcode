package p316

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, "acdb", removeDuplicateLetters("cbacdcbc"))
}
