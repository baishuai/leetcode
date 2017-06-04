package p357

import "testing"
import "github.com/stretchr/testify/assert"

func Test0(t *testing.T) {
	assert.Equal(t, 1, countNumbersWithUniqueDigits(0))
	assert.Equal(t, 10, countNumbersWithUniqueDigits(1))
	assert.Equal(t, 91, countNumbersWithUniqueDigits(2))
	assert.Equal(t, 739, countNumbersWithUniqueDigits(3))
}

func Test1(t *testing.T) {
	assert.Equal(t, 8877691, countNumbersWithUniqueDigits(15))
	assert.Equal(t, 8877691, countNumbersWithUniqueDigits(16))
}
