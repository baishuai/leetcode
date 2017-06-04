package p330

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, minPatches([]int{1, 3}, 6))
}

func Test1(t *testing.T) {
	assert.Equal(t, 2, minPatches([]int{1, 5, 10}, 20))
}

func Test2(t *testing.T) {
	assert.Equal(t, 0, minPatches([]int{1, 2, 2}, 5))
}
