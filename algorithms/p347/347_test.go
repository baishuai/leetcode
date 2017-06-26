package p347

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, []int{1, 2}, topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	assert.Equal(t, []int{1, 2}, topKFrequent([]int{5, 4, 1, 1, 1, 2, 2, 3}, 2))
}
