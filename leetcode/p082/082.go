package p082

/**
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

For example,
Given 1->2->3->3->4->4->5, return 1->2->5.
Given 1->1->1->2->3, return 2->3.

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

func deleteDuplicates(head *ListNode) *ListNode {

	head = &ListNode{0, head}
	quick, slow := head.Next, head
	for quick != nil {
		if quick.Next != nil && quick.Next.Val == quick.Val {
			dupv := quick.Val
			for quick != nil && quick.Val == dupv {
				quick = quick.Next
			}
		} else {
			slow.Next = quick
			slow = slow.Next
			quick = quick.Next
		}
	}
	slow.Next = quick
	return head.Next
}
