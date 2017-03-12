package p086

/**
Given a linked list and a value x, partition it such that all nodes less than x
 come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

For example,
Given 1->4->3->2->5->2 and x = 3,
return 1->2->2->4->3->5.
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

func partition(head *ListNode, x int) *ListNode {
	master := &ListNode{0, head}

	slow, quick := master, master
	for quick.Next != nil {
		if quick.Next.Val < x {
			if slow != quick {
				tmp := quick.Next
				quick.Next = quick.Next.Next
				tmp.Next = slow.Next
				slow.Next = tmp
			} else {
				quick = quick.Next
			}
			slow = slow.Next
		} else {
			quick = quick.Next
		}
	}

	return master.Next
}
