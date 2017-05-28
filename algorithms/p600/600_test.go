package p600

import "testing"
import "github.com/stretchr/testify/assert"

func Test0(t *testing.T) {
	assert.Equal(t, 5, findIntegers(5))
}

func Test1(t *testing.T) {
	assert.Equal(t, 2, findIntegers(1))
}

func Test2(t *testing.T) {
	assert.Equal(t, 4, findIntegers(4))
}
