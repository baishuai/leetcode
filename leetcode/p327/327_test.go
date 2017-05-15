package p327

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 3, countRangeSum([]int{-2, 5, -1}, -2, 2))
}
