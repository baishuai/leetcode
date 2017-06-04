package p446

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 7, numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
}

func Test1(t *testing.T) {
	assert.Equal(t, 13, numberOfArithmeticSlices([]int{2, 4, 6, 8, 10, 11, 12}))
}

func Test2(t *testing.T) {
	assert.Equal(t, 24, numberOfArithmeticSlices([]int{2, 4, 6, 8, 10, 10, 10, 10}))
}

func Test3(t *testing.T) {
	assert.Equal(t, 2, numberOfArithmeticSlices([]int{2, 2, 3, 4}))
}

func Test4(t *testing.T) {
	assert.Equal(t, 0, numberOfArithmeticSlices([]int{}))
	assert.Equal(t, 0, numberOfArithmeticSlices([]int{1}))
	assert.Equal(t, 0, numberOfArithmeticSlices([]int{1, 2}))
}
