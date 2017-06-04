package p321

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, []int{5, 4}, maxOneArray([]int{5, 2, 3, 4, 1}, 2))
	assert.Equal(t, []int{6}, maxOneArray([]int{2, 4, 6, 5}, 1))
	assert.Equal(t, []int{6, 5}, maxOneArray([]int{2, 4, 6, 5}, 2))
	assert.Equal(t, []int{4, 6, 5}, maxOneArray([]int{2, 4, 6, 5}, 3))
	assert.Equal(t, []int{2, 4, 6, 5}, maxOneArray([]int{2, 4, 6, 5}, 4))
}

func Test1(t *testing.T) {
	assert.Equal(t, []int{9, 8, 6, 5, 3}, maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
}
