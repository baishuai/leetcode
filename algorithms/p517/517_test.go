package p517

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 3, findMinMoves([]int{1, 0, 5}))
}

func Test1(t *testing.T) {
	assert.Equal(t, 2, findMinMoves([]int{0, 3, 0}))

}

func Test2(t *testing.T) {
	assert.Equal(t, -1, findMinMoves([]int{0, 2, 0}))
}

func Test3(t *testing.T) {
	assert.Equal(t, 4, findMinMoves([]int{9, 1, 8, 8, 9}))
}
