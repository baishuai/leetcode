package p368

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, []int{1, 2}, largestDivisibleSubset([]int{1, 2, 3}))
}

func Test1(t *testing.T) {
	assert.Equal(t, []int{}, largestDivisibleSubset([]int{}))
}
