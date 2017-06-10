package p324

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func sign(diff int) int {
	if diff < 0 {
		return -1
	} else if diff > 0 {
		return 1
	}
	return 0
}

func test(nums []int) bool {
	diffSign := 1
	for i := 1; i < len(nums); i++ {
		if sign(nums[i]-nums[i-1]) != diffSign {
			return false
		}
		diffSign = 0 - diffSign
	}
	return true
}

func Test0(t *testing.T) {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	assert.True(t, test(nums))
}
