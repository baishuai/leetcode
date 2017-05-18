package p410

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, splitArrayMax([]int64{1, 2, 3}, 3))
	assert.Equal(t, 2, splitArrayMax([]int64{1, 2, 3}, 2))
}

func Test1(t *testing.T) {
	assert.Equal(t, 18, splitArray([]int{7, 2, 5, 10, 8}, 2))
}

func Test2(t *testing.T) {
	assert.Equal(t, 2, splitArray([]int{1, 1}, 1))
	assert.Equal(t, 1, splitArray([]int{1}, 1))
	assert.Equal(t, 1, splitArray([]int{1, 1}, 2))
}

func Test3(t *testing.T) {
	assert.Equal(t, 2147483647, splitArray([]int{1, 2147483647}, 2))
}

func Test4(t *testing.T) {
	assert.Equal(t, 6, splitArray([]int{5, 2, 4, 1, 3, 6, 0}, 4))
}
