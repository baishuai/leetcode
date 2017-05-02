package p188

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 0, maxProfit(1, []int{}))
	assert.Equal(t, 0, maxProfit(1, []int{100}))
}

func Test1(t *testing.T) {
	assert.Equal(t, 8773, maxProfit(7, []int{2, 2, 3, 345, 5, 7667, 21, 32, 33, 5, 234, 5, 5, 214, 532, 34, 22, 2}))
}
