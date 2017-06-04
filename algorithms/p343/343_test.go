package p343

import "testing"
import "github.com/stretchr/testify/assert"

func Test0(t *testing.T) {
	assert.Equal(t, 36, integerBreak(10))
}

func Test2(t *testing.T) {
	assert.Equal(t, 1, integerBreak(2))
}
