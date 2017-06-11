package p611

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 3, triangleNumber([]int{2, 2, 3, 4}))
}
