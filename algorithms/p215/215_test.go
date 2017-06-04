package p215

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 4, findKthLargest([]int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 7))
}
