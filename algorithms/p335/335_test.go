package p335

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, true, isSelfCrossing([]int{1, 1, 1, 1}))
}

func Test1(t *testing.T) {
	assert.Equal(t, true, isSelfCrossing([]int{2, 1, 1, 2}))
}

func Test2(t *testing.T) {
	assert.Equal(t, false, isSelfCrossing([]int{1, 2, 3, 4}))
}

func Test3(t *testing.T) {
	assert.Equal(t, true, isSelfCrossing([]int{2, 2, 3, 3, 2, 2}))
}

func Test4(t *testing.T) {
	assert.Equal(t, true, isSelfCrossing([]int{2, 2, 3, 2, 1}))
}
