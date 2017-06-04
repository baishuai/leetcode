package p143

/**
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You must do this in-place without altering the nodes' values.

For example,
Given {1,2,3,4}, reorder it to {1,4,2,3}.
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

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	master := &ListNode{0, head}
	slow, quick := master, head
	for quick != nil {
		slow = slow.Next
		quick = quick.Next
		if quick != nil {
			quick = quick.Next
		}
	}

	quick = slow.Next
	slow.Next = nil

	cur := quick.Next
	quick.Next = nil
	for cur != nil {
		tmp := cur
		cur = cur.Next
		tmp.Next = quick
		quick = tmp
	}

	slow = head
	for quick != nil {
		tmp := quick
		quick = quick.Next
		tmp.Next = slow.Next
		slow.Next = tmp
		slow = tmp.Next
	}
}
