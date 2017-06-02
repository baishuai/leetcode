package p264

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, nthUglyNumber(1))
	assert.Equal(t, 2, nthUglyNumber(2))
	assert.Equal(t, 3, nthUglyNumber(3))
	assert.Equal(t, 4, nthUglyNumber(4))
}

func Test1(t *testing.T) {
	assert.Equal(t, 1536, nthUglyNumber(100))
}
