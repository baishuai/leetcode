package p441

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, arrangeCoins(1))
	assert.Equal(t, 2, arrangeCoins(3))
	assert.Equal(t, 2, arrangeCoins(4))
	assert.Equal(t, 2, arrangeCoins(5))
	assert.Equal(t, 3, arrangeCoins(6))
}
