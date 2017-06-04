package p132

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, minCut("aab"))
	assert.Equal(t, 1, minCut2("aab"))
}
