package p304

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	mat := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	nm := Constructor(mat)
	assert.Equal(t, 8, nm.SumRegion(2, 1, 4, 3))
	assert.Equal(t, 11, nm.SumRegion(1, 1, 2, 2))
	assert.Equal(t, 12, nm.SumRegion(1, 2, 2, 4))
}

func Test1(t *testing.T) {
	mat := [][]int{}
	Constructor(mat)
}

func Test2(t *testing.T) {
	mat := [][]int{[]int{}}
	Constructor(mat)
}
