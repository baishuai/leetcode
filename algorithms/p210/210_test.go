package p210

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	ans := findOrder(2, [][]int{{1, 0}})

	assert.Equal(t, []int{0, 1}, ans)
}

func Test1(t *testing.T) {
	ans := findOrder(2, [][]int{{1, 0}, {0, 1}})
	assert.Equal(t, []int{}, ans)
}
