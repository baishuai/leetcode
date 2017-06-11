package p365

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGcd(t *testing.T) {
	assert.Equal(t, 3, gcd(3, 6))
	assert.Equal(t, 3, gcd(9, 6))
	assert.Equal(t, 0, gcd(0, 0))
	assert.Equal(t, 2, gcd(2, 0))
}

func TestWater(t *testing.T) {
	assert.Equal(t, true, canMeasureWater(3, 5, 4))
	assert.Equal(t, false, canMeasureWater(2, 6, 5))
	assert.Equal(t, true, canMeasureWater(0, 0, 0))
}
