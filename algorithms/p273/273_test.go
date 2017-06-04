package p273

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, "One Hundred Twenty Three", numberToWords(123))
}

func Test1(t *testing.T) {
	assert.Equal(t, "Twelve Thousand Three Hundred Forty Five", numberToWords(12345))
}

func Test2(t *testing.T) {
	assert.Equal(t, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven", numberToWords(1234567))
}

func Test3(t *testing.T) {
	assert.Equal(t, "One Hundred", numberToWords(100))
}
