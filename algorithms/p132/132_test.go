package p132

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, minCut("aab"))
	assert.Equal(t, 1, minCut2("aab"))
}
