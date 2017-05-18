package p316

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, "acdb", removeDuplicateLetters("cbacdcbc"))
}
