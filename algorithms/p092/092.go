package p092

/**
Reverse a linked list from position m to n. Do it in-place and in one-pass.

For example:
Given 1->2->3->4->5->NULL, m = 2 and n = 4,

return 1->4->3->2->5->NULL.

Note:
Given m, n satisfy the following condition:
1 ≤ m ≤ n ≤ length of list.
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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	master := &ListNode{0, head}
	slow := master
	n = n - m
	for m > 1 {
		slow = slow.Next
		m--
	}
	// slow 指向 第m个节点的前一个节点
	quick := slow
	for n >= 0 {
		quick = quick.Next
		n--
	}
	// quick 刚好是第n个节点
	front := slow
	slow = slow.Next

	next := slow.Next
	slow.Next = quick.Next
	for slow != quick {
		next.Next, slow, next = slow, next, next.Next
	}
	front.Next = slow

	return master.Next
}
