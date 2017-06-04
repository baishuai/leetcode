package p576

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 12, findPaths(1, 3, 3, 0, 1))
}

func Test1(t *testing.T) {
	assert.Equal(t, 6, findPaths(2, 2, 2, 0, 0))
}
