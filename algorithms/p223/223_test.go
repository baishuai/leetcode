package p223

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 45, computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
	assert.Equal(t, 33, computeArea(-3, 0, 3, 4, 0, -1, 9, 0))
}
