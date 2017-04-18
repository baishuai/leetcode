package p322

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	ans := coinChange([]int{1, 2, 5}, 11)
	assert.Equal(t, 3, ans)
}

func Test1(t *testing.T) {
	ans := coinChange([]int{1}, 0)
	assert.Equal(t, 0, ans)
}
