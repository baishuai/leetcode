package p219

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, false, containsNearbyDuplicate([]int{1, 2, 3, 4, 5}, 3))
	assert.Equal(t, true, containsNearbyDuplicate([]int{1, 2, 4, 4, 5}, 1))
	assert.Equal(t, false, containsNearbyDuplicate([]int{}, 1))
}
