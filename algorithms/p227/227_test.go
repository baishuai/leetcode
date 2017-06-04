package p227

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 7, calculate("3+2*2"))
}

func Test1(t *testing.T) {
	assert.Equal(t, 1, calculate(" 3/2 "))
}

func Test2(t *testing.T) {
	assert.Equal(t, 5, calculate(" 3+5 / 2 "))
}

func Test3(t *testing.T) {
	assert.Equal(t, 3, calculate("1+1+1"))
}

func Test4(t *testing.T) {
	assert.Equal(t, 24, calculate("2*3*4"))
}
