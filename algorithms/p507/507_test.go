package p507

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, true, checkPerfectNumber(28))
	assert.Equal(t, false, checkPerfectNumber(27))
	assert.Equal(t, false, checkPerfectNumber(1))
}
