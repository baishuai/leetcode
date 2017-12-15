package p724

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	assert.Equal(t, 3, pivotIndex(nums))
}

func Test1(t *testing.T) {
	nums := []int{1, 2, 3}
	assert.Equal(t, -1, pivotIndex(nums))
}
