package p028

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, -1, strStr("bai", "ij"))
}

func Test1(t *testing.T) {
	assert.Equal(t, 1, strStr("bai", "ai"))
}