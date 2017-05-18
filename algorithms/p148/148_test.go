package p148

import "testing"

func test(t *testing.T, head *ListNode, count int) {
	head = sortList(head)

	if head == nil {
		return
	}
	last := head.Val

	head = head.Next
	for head != nil {
		if head.Val < last {
			t.Fatal("error")
		} else {
			last = head.Val
		}
		count--
		head = head.Next
	}
	if count != 1 {
		t.Fatal("error length")
	}
}

func TestName(t *testing.T) {
	test(t, &ListNode{1, &ListNode{2, nil}}, 2)
}
