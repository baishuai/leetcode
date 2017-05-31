package p416

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, true, canPartition([]int{1, 5, 11, 5}))
}
