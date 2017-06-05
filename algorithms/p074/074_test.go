package p074

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	mat := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	assert.Equal(t, true, searchMatrix(mat, 3))
	assert.Equal(t, false, searchMatrix(mat, 4))
}

func Test1(t *testing.T) {
	mat := [][]int{}
	assert.Equal(t, false, searchMatrix(mat, 4))
}
