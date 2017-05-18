package p203

import "testing"

func TestExample(t *testing.T) {
	head := removeElements(&ListNode{Val: 0}, 0)
	if head != nil {
		t.FailNow()
	}
}
