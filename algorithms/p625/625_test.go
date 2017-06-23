package p625

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, smallestFactorization(1))
}

func Test2(t *testing.T) {
	assert.Equal(t, 68, smallestFactorization(48))
	assert.Equal(t, 0, smallestFactorization(67))
	assert.Equal(t, 0, smallestFactorization(1220703125))
}
