package p028

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, -1, strStr("bai", "ij"))
}

func Test1(t *testing.T) {
	assert.Equal(t, 1, strStr("bai", "ai"))
}
