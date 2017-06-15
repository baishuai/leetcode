package p274

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 3, hIndex([]int{3, 0, 6, 1, 5}))
}
