package p239

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	ans := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, ans)
}

func Test1(t *testing.T) {
	ans := maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3)
	assert.Equal(t, []int{3, 3, 2, 5}, ans)
}
