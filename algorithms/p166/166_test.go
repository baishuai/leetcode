package p166

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, "0.5", fractionToDecimal(1, 2))
}

func Test1(t *testing.T) {
	assert.Equal(t, "2", fractionToDecimal(2, 1))
}

func Test2(t *testing.T) {
	assert.Equal(t, "0.(6)", fractionToDecimal(2, 3))
}

func Test3(t *testing.T) {
	assert.Equal(t, "-6.25", fractionToDecimal(-50, 8))
}
