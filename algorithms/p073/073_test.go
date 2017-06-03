package p073

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	mat := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}

	setZeroes(mat)
	exp := [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	assert.Equal(t, exp, mat)
}


func Test1(t *testing.T) {
	mat := [][]int{
		{1, 1, 1},
		{0, 0, 1},
		{1, 1, 0},
	}

	setZeroes(mat)
	exp := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	assert.Equal(t, exp, mat)
}


func Test2(t *testing.T) {
	mat := [][]int{
		{1, 0, 1},
		{1, 0, 1},
		{1, 1, 0},
	}

	setZeroes(mat)
	exp := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	assert.Equal(t, exp, mat)
}


func Test3(t *testing.T) {
	mat := [][]int{

	}

	setZeroes(mat)
	exp := [][]int{
	}

	assert.Equal(t, exp, mat)
}

