package p023

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	t1 := &ListNode{Val: 1}
	t2 := &ListNode{Val: 2}
	t3 := &ListNode{Val: 3}
	t4 := &ListNode{Val: 4}
	t5 := &ListNode{Val: 5}

	lists := []*ListNode{t1, t2, t3, t4, t5}

	list := mergeKLists(lists)

	assert.True(t, list != nil)
	i := 1
	for list != nil {
		assert.Equal(t, i, list.Val)
		i++
		list = list.Next
	}
}
