package p326

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	i := 1
	for i < (1 << 30) {
		i *= 3
		assert.Equal(t, true, isPowerOfThree(i))
		assert.Equal(t, false, isPowerOfThree(i+1))
	}
}
