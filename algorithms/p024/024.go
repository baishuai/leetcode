package p024

/**
Given a linked list, swap every two adjacent nodes and return its head.

For example,
Given 1->2->3->4, you should return the list as 2->1->4->3.

Your algorithm should use only constant space.
You may not modify the values in the list, only nodes itself can be changed.
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

func swapPairs(head *ListNode) *ListNode {
	head = &ListNode{0, head}

	var cur, left, right *ListNode
	cur = head
	for cur.Next != nil && cur.Next.Next != nil {
		left = cur.Next
		right = cur.Next.Next

		cur.Next = right
		left.Next = right.Next
		right.Next = left

		cur = left
	}
	return head.Next
}
