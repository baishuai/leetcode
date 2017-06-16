package p400

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, findNthDigit(1))
	assert.Equal(t, 0, findNthDigit(11))
}
