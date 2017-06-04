package p233

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, countDigitOne(3))
}

func Test1(t *testing.T) {
	assert.Equal(t, 6, countDigitOne(13))
}

func Test2(t *testing.T) {
	assert.Equal(t, 161, countDigitOne(301))
}

func Test3(t *testing.T) {
	assert.Equal(t, 0, countDigitOne(-1))
	assert.Equal(t, 0, countDigitOne(0))
}
