package p174

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	du := [][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}

	assert.Equal(t, 3, calculateMinimumHP(du))
}

func Test1(t *testing.T) {
	du := [][]int{{10}}
	assert.Equal(t, 1, calculateMinimumHP(du))
}
