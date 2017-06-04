package p021

/**
Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of the first two lists.
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	head := &ListNode{0, l1}

	l1 = head
	for l1 != nil && l2 != nil {
		for l1.Next != nil && l1.Next.Val <= l2.Val {
			l1 = l1.Next
		}
		l1.Next, l2 = l2, l1.Next
	}
	return head.Next
}
