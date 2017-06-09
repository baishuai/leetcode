package p225

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	assert.Equal(t, 1, obj.Top())
	assert.Equal(t, 1, obj.Pop())
	assert.Equal(t, true, obj.Empty())

	assert.Equal(t, -1, obj.Top())
	assert.Equal(t, -1, obj.Pop())
}
