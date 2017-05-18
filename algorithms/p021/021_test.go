package p021

import (
	"testing"
)

func TestContact(t *testing.T) {
	l1 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, &ListNode{9, nil}}}}}
	l2 := &ListNode{2, &ListNode{4, &ListNode{6, &ListNode{8, nil}}}}

	ans := mergeTwoLists(l1, l2)
	index := 1
	for ans != nil {
		if ans.Val != index {
			t.Fatalf("error answer, %v", ans)
		}
		index ++
		ans = ans.Next
	}
}

func TestSame(t *testing.T) {
	l1 := &ListNode{1, nil }
	l2 := &ListNode{1, nil}

	ans := mergeTwoLists(l1, l2)
	count := 0
	for ans != nil {
		if ans.Val != 1 {
			t.Fatalf("error answer, %v", ans)
		}
		count ++
		ans = ans.Next
	}
	if count != 2 {
		t.Fatalf("error answer, count %v", count)
	}
}
