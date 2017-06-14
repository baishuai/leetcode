package p231

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, false, isPowerOfTwo(0))
}

func Test1(t *testing.T) {
	for i := uint(1); i < 32; i++ {
		assert.Equal(t, true, isPowerOfTwo(1<<i))
		assert.Equal(t, false, isPowerOfTwo((1<<i)+1))
	}
}
