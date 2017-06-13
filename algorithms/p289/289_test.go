package p289

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	board := [][]int{
		{1, 0, 0, 0, 0, 1},
		{0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1, 0},
		{1, 1, 1, 1, 0, 1},
		{0, 1, 1, 0, 1, 0},
		{1, 0, 1, 0, 1, 1},
		{1, 0, 0, 1, 1, 1},
		{1, 1, 0, 0, 0, 0},
	}

	exp := [][]int{
		{0, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 1},
		{1, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0},
		{1, 0, 1, 1, 0, 1},
		{1, 1, 0, 0, 1, 0},
	}
	gameOfLife(board)

	assert.Equal(t, exp, board)
}

func Test2(t *testing.T) {
	board := [][]int{}
	gameOfLife(board)
	assert.Equal(t, [][]int{}, board)
}
