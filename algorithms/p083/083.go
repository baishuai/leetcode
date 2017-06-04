package p083

/**
Given a sorted linked list, delete all duplicates such that
 each element appear only once.

For example,
Given 1->1->2, return 1->2.
Given 1->1->2->3->3, return 1->2->3.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	slow := head
	qv := head.Val
	for quick := head.Next; quick != nil; quick = quick.Next {
		if quick.Val == qv {
			continue
		} else {
			qv = quick.Val
			slow.Next = quick
			slow = slow.Next
		}
	}
	slow.Next = nil
	return head
}
