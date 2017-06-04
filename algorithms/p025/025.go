package p025

/**
Given a linked list, reverse the nodes of a linked list k at a time
 and return its modified list.

k is a positive integer and is less than or equal to the length
 of the linked list. If the number of nodes is not a multiple
 of k then left-out nodes in the end should remain as it is.

You may not alter the values in the nodes, only nodes itself may be changed.

Only constant memory is allowed.

For example,
Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

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

func reverseKGroup(head *ListNode, k int) *ListNode {

	head = &ListNode{0, head}
	quick, slow := head.Next, head
	for quick != nil {
		count := 0
		for quick != nil {
			quick = quick.Next
			if count++; count == k {
				break
			}
		}
		if count == k {
			lastEnd := slow
			slow = slow.Next
			reverse := lastEnd.Next.Next
			for count > 1 {
				prevNext := lastEnd.Next
				lastEnd.Next = reverse
				reverse = reverse.Next
				lastEnd.Next.Next = prevNext
				count--
			}
			slow.Next = quick
		}
	}

	return head.Next
}
