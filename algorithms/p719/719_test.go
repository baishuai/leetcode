package p719

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {

	nums := []int{1, 3, 1}
	assert.Equal(t, 0, smallestDistancePair(nums, 1))
}

func TestCount(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 7, countDistance(nums, 2))
}
