package p581

import "testing"
import "github.com/stretchr/testify/assert"

func Test0(t *testing.T) {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	assert.Equal(t, 5, findUnsortedSubarray(nums))
}

func Test1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	assert.Equal(t, 0, findUnsortedSubarray(nums))
}

func Test2(t *testing.T) {
	nums := []int{1, 3, 5, 2, 4}
	assert.Equal(t, 4, findUnsortedSubarray(nums))
}
