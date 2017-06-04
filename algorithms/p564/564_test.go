package p564

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, "12321", nearestPalindromic("12345"))
}

func Test1(t *testing.T) {
	assert.Equal(t, "123321", nearestPalindromic("123345"))
}

func Test2(t *testing.T) {
	assert.Equal(t, "99", nearestPalindromic("100"))
}

func Test3(t *testing.T) {
	assert.Equal(t, "1233321", nearestPalindromic("1234321"))
}

func Test4(t *testing.T) {
	assert.Equal(t, "1001", nearestPalindromic("999"))
	assert.Equal(t, "10001", nearestPalindromic("9999"))
}

func Test5(t *testing.T) {
	assert.Equal(t, "9", nearestPalindromic("11"))
	assert.Equal(t, "11", nearestPalindromic("22"))
	assert.Equal(t, "22", nearestPalindromic("23"))
}

func Test6(t *testing.T) {
	assert.Equal(t, "101", nearestPalindromic("99"))
}
