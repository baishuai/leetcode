package p263

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, false, isUgly(0))
	assert.Equal(t, true, isUgly(1))
	assert.Equal(t, true, isUgly(30))
}
