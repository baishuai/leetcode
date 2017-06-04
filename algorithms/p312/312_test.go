package p312

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 63, maxCoins([]int{2, 9, 3}))
}

func Test1(t *testing.T) {
	assert.Equal(t, 84, maxCoins([]int{2, 9, 4}))
}

func Test2(t *testing.T) {
	assert.Equal(t, 167, maxCoins([]int{3, 1, 5, 8}))
}

func Test3(t *testing.T) {
	assert.Equal(t, 216, maxCoins([]int{3, 4, 5, 6}))
}

func Test4(t *testing.T) {
	assert.Equal(t, 277, maxCoins([]int{3, 4, 5, 5, 4, 3}))
}

func Test5(t *testing.T) {
	assert.Equal(t, 0, maxCoins([]int{}))
}

func Test6(t *testing.T) {
	assert.Equal(t, 6, maxCoins([]int{6}))
}

func Test7(t *testing.T) {
	assert.Equal(t, 6, maxCoins([]int{2, 2}))
}

func Test8(t *testing.T) {
	assert.Equal(t, 1088290, maxCoins([]int{9, 76, 64, 21, 97, 60, 5}))
}

func Test9(t *testing.T) {
	assert.Equal(t, 10, maxCoins([]int{1, 5}))
}

func Test10(t *testing.T) {
	assert.Equal(t, 1849648, maxCoins([]int{35, 16, 83, 87, 84, 59, 48, 41, 20, 54}))
}

func Test11(t *testing.T) {
	assert.Equal(t, 1780, maxCoins([]int{7, 9, 8, 0, 7, 1, 3, 5, 5, 2, 3, 3, 3}))
}

func Test12(t *testing.T) {
	assert.Equal(t, 8118, maxCoins([]int{0, 0, 4, 8, 0, 2, 4, 2, 4, 4, 5, 9, 2, 5, 4, 6, 9, 8, 2, 4, 4, 3, 6, 4, 8, 0, 4, 4, 5, 8, 7, 6, 0, 6, 6, 8}))
}
