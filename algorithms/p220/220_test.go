package p220

import "testing"
import "github.com/stretchr/testify/assert"

func TestLess2(t *testing.T) {
	assert.Equal(t, false, containsNearbyAlmostDuplicate([]int{}, 0, 0))
	assert.Equal(t, false, containsNearbyAlmostDuplicate([]int{0}, 0, 0))
}

func Test2(t *testing.T) {
	assert.Equal(t, false, containsNearbyAlmostDuplicate([]int{-1, -1}, 0, 0))
}

func Test3(t *testing.T) {
	assert.Equal(t, true, containsNearbyAlmostDuplicate([]int{3, 6, 0, 4}, 2, 2))
}

func Test4(t *testing.T) {
	assert.Equal(t, false, containsNearbyAlmostDuplicate([]int{-3, 3}, 2, 4))
}
