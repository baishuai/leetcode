package p268

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 2, missingNumber([]int{0, 1, 3, 4}))
}
