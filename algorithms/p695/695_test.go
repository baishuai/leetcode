package p695

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	grid := [][]int{
		{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1},
	}
	assert.Equal(t, 4, maxAreaOfIsland(grid))
}
