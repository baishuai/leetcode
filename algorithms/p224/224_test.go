package p224

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 2, calculate("1 + 1"))
}

func Test1(t *testing.T) {
	assert.Equal(t, 3, calculate(" 2-1 + 2 "))
}

func Test2(t *testing.T) {
	assert.Equal(t, 23, calculate("(1+(4+5+2)-3)+(6+8)"))
}

func Test3(t *testing.T) {
	assert.Equal(t, 12345, calculate("12345"))
}

func Test4(t *testing.T) {
	assert.Equal(t, 99, calculate("    99  "))
}
