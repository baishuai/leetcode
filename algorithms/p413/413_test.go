package p413

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 6, numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
}

func Test1(t *testing.T) {
	assert.Equal(t, 3, numberOfArithmeticSlices([]int{2, 4, 6, 8}))
}

func Test2(t *testing.T) {
	assert.Equal(t, 10, numberOfArithmeticSlices([]int{1, 2, 3, 4, 5, 6}))
}
