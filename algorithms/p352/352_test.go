package p352

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	obj := Constructor()
	obj.Addnum(1)
	obj.Addnum(3)
	obj.Addnum(7)
	obj.Addnum(2)
	obj.Addnum(6)
	intervals := obj.Getintervals()
	assert.Equal(t, intervals, []Interval{{1, 3}, {6, 7}})
}
