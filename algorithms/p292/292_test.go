package p292

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {

	assert.True(t, canWinNim(1))
	assert.True(t, canWinNim(2))
	assert.True(t, canWinNim(3))
	assert.True(t, canWinNim(5))
	assert.False(t, canWinNim(4))
}
