package p576

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 12, findPaths(1, 3, 3, 0, 1))
}

func Test1(t *testing.T) {
	assert.Equal(t, 6, findPaths(2, 2, 2, 0, 0))
}
