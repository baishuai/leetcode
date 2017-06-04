package p605

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, true, canPlaceFlowers([]int{0}, 1))
}

func Test1(t *testing.T) {
	assert.Equal(t, false, canPlaceFlowers([]int{0, 1, 0}, 1))
}

func Test2(t *testing.T) {
	assert.Equal(t, false, canPlaceFlowers([]int{0, 1}, 1))
}

func Test3(t *testing.T) {
	assert.Equal(t, true, canPlaceFlowers([]int{0, 0}, 1))
}

func Test4(t *testing.T) {
	assert.Equal(t, true, canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}

func Test5(t *testing.T) {
	assert.Equal(t, true, canPlaceFlowers([]int{0, 0, 1, 0, 0, 0, 1, 0, 0}, 1))
}
