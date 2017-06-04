package p494

import "testing"
import "github.com/stretchr/testify/assert"

func Test0(t *testing.T) {
	assert.Equal(t, 5, findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func Test1(t *testing.T) {
	assert.Equal(t, 10, findTargetSumWays([]int{1, 1, 1, 1, 1, 0}, 3))
}
