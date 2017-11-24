package p689

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	nums := []int{1, 2, 1, 2, 6, 7, 5, 1}
	exp := []int{0, 3, 5}
	assert.Equal(t, exp, maxSumOfThreeSubarrays(nums, 2))
}

func Test1(t *testing.T) {

	nums := []int{7, 13, 20, 19, 19, 2, 10, 1, 1, 19}
	exp := []int{1, 4, 7}
	assert.Equal(t, exp, maxSumOfThreeSubarrays(nums, 3))
}
