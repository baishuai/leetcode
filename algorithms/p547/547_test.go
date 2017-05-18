package p547

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	ans := findCircleNum([][]int{
		{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
	})

	assert.Equal(t, 2, ans)
}
