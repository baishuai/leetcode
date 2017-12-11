package p685

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {

	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}
	assert.Equal(t, []int{4, 1}, findRedundantDirectedConnection(edges))
}
