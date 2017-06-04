package p203

/**
Remove all elements from a linked list of integers that have value val.

Example
Given: 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6, val = 6
Return: 1 --> 2 --> 3 --> 4 --> 5
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

func removeElements(head *ListNode, val int) *ListNode {
	master := &ListNode{0, head}
	quick, slow := head, master
	for quick != nil {
		if quick.Val == val {
			slow.Next = quick.Next
			quick = quick.Next
		} else {
			quick = quick.Next
			slow = slow.Next
		}
	}
	return master.Next
}
