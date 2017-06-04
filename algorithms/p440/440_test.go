package p440

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 10, findKthNumber(13, 2))
}

func Test1(t *testing.T) {
	assert.Equal(t, 3, findKthNumber(302, 223))
}

func Test2(t *testing.T) {
	assert.Equal(t, 3008, findKthNumber(3022, 2234))
}
