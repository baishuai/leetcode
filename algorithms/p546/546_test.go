package p546

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 23, removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}))
}

func Test2(t *testing.T) {
	assert.Equal(t, 25, removeBoxes([]int{1, 1, 1, 1, 1}))
}
