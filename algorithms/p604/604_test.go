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
	t.Log(iterator.Next())                      // return 'C'
	t.Log(iterator.Next())                      // return 'o'
	t.Log(iterator.Next())                      // return 'd'
	t.Log(iterator.HasNext())                   // return true
	t.Log(iterator.Next())                      // return 'e'
	t.Log(iterator.HasNext())                   // return false
	t.Log(iterator.Next())                      // return ' '
}
