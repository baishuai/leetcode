package p403

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.True(t, canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
}

func Test1(t *testing.T) {
	assert.False(t, canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
}
