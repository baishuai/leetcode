package p539

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 61, minutes("01:01"))
}

func Test0(t *testing.T) {
	assert.Equal(t, 1, findMinDifference([]string{"23:59", "00:00"}))
	assert.Equal(t, 1, findMinDifference([]string{"00:00", "23:59", "20:50"}))
}
