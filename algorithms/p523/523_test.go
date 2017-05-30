package p523

import "testing"
import "github.com/stretchr/testify/assert"

func Test0(t *testing.T) {
	assert.Equal(t, true, checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
}

func Test1(t *testing.T) {
	assert.Equal(t, false, checkSubarraySum([]int{6}, 6))
}

func Test2(t *testing.T) {
	assert.Equal(t, true, checkSubarraySum([]int{6, 0}, 6))
}

func Test3(t *testing.T) {
	assert.Equal(t, false, checkSubarraySum([]int{23, 2, 4, 6, 7}, 0))
}
