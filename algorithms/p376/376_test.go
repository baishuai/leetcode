package p376

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 2, wiggleMaxLength([]int{1, 2, 3, 4, 5, 6}))
	assert.Equal(t, 6, wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	assert.Equal(t, 7, wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	assert.Equal(t, 1, wiggleMaxLength([]int{1, 1, 1}))
	assert.Equal(t, 0, wiggleMaxLength([]int{}))
}
