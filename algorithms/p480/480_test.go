package p480

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, []float64{1, -1, -1, 3, 5, 6}, medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
