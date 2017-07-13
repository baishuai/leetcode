package p621

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 8, leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}

func Test1(t *testing.T) {
	assert.Equal(t, 5, leastInterval([]byte{'A', 'A', 'B', 'B', 'Z'}, 1))
}
