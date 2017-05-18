package p019

import "testing"

func TestOne(t *testing.T) {
	head := &ListNode{Val: 1, Next: nil}

	if removeNthFromEnd(head, 1) != nil {
		t.Fatal("error answer")
	}

}

func TestTwo(t *testing.T) {
	head := &ListNode{1, &ListNode{2, nil}}

	if ans := removeNthFromEnd(head, 2) ; ans.Val != 2 || ans.Next != nil {
		t.Fatal("error answer")
	}

}

func TestTwo2(t *testing.T) {
	head := &ListNode{1, &ListNode{2, nil}}

	if ans := removeNthFromEnd(head, 1) ; ans.Val != 1 || ans.Next != nil {
		t.Fatal("error answer")
	}

}