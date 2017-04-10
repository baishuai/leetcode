package p128

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
