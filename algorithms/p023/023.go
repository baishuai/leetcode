package p023

/**
Merge k sorted linked lists and return it as one sorted list.
 Analyze and describe its complexity.
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	} else {
		half := len(lists) / 2
		return mergeTwoLists(mergeKLists(lists[:half]), mergeKLists(lists[half:]))
	}
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
