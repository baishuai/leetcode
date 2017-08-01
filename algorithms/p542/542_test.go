package p542

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test2(t *testing.T) {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	dis := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 2, 1},
	}

	assert.Equal(t, dis, updateMatrix(mat))
}

func Test1(t *testing.T) {
	assert.Nil(t, updateMatrix(nil))
}
