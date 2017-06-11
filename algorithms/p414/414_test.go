package p414

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, thirdMax([]int{3, 2, 1}))
	assert.Equal(t, 2, thirdMax([]int{1, 2}))
	assert.Equal(t, 1, thirdMax([]int{2, 2, 1, 3}))
}
