package p327

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 3, countRangeSum([]int{-2, 5, -1}, -2, 2))
}
