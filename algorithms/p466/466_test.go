package p466

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 2, getMaxRepetitions("acb", 4, "ab", 2))
}

func Test1(t *testing.T) {
	assert.Equal(t, 10, getMaxRepetitions("a", 1000000, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 1000))
}
