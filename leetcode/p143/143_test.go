package p143

import "testing"

func TestEXample(t *testing.T) {
	reorderList(&ListNode{1, &ListNode{2, &ListNode{3, nil}}})
	reorderList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}})
	reorderList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}})
}
