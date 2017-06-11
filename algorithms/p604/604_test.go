package p604

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	iterator := Constructor("L1e2t1C1o1d1e1")

	assert.Equal(t, byte('L'), iterator.Next()) // return 'L'
	assert.Equal(t, byte('e'), iterator.Next()) // return 'e'
	assert.Equal(t, byte('e'), iterator.Next()) // return 'e'
	assert.Equal(t, byte('t'), iterator.Next()) // return 't'
	assert.Equal(t, byte('C'), iterator.Next()) // return 'C'
	assert.Equal(t, byte('o'), iterator.Next()) // return 'o'
	assert.Equal(t, byte('d'), iterator.Next()) // return 'd'
	assert.Equal(t, true, iterator.HasNext())   // return true
	assert.Equal(t, byte('e'), iterator.Next()) // return 'e'
	assert.Equal(t, false, iterator.HasNext())  // return false
	assert.Equal(t, byte(' '), iterator.Next()) // return ' '
}
