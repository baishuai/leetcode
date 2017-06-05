package p392

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, true, isSubsequence("abc", "ahbgcd"))
	assert.Equal(t, false, isSubsequence("axc", "ahbgcd"))
}
