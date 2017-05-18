package p061


/**
Given a list, rotate the list to the right by k places, where k is non-negative.

For example:
Given 1->2->3->4->5->NULL and k = 2,
return 4->5->1->2->3->NULL.
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	newHead := &ListNode{0, head}
	slow, quick := newHead, newHead
	listLen := 0
	for quick.Next != nil {
		quick = quick.Next
		listLen++
	}
	if listLen == 0 {
		return newHead.Next
	}
	k = listLen - k%listLen
	for k > 0 {
		slow = slow.Next
		k--
	}
	quick.Next = newHead.Next
	newHead.Next = slow.Next
	slow.Next = nil
	return newHead.Next
}
