package p467

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 6, findSubstringInWraproundString("zab"))
}

func Test1(t *testing.T) {
	assert.Equal(t, 1, findSubstringInWraproundString("a"))
}

func Test2(t *testing.T) {
	assert.Equal(t, 10, findSubstringInWraproundString("czabcc"))
}
