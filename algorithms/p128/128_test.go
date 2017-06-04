package p128

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
