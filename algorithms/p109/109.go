package p109

/**
Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {

	if head == nil {
		return nil
	} else if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	head = &ListNode{0, head}
	slow, quick := head, head.Next
	for quick != nil {
		quick = quick.Next
		if quick != nil {
			slow = slow.Next
			quick = quick.Next
		}
	}

	quick = slow.Next
	slow.Next = nil
	node := &TreeNode{Val: quick.Val}
	node.Left = sortedListToBST(head.Next)
	node.Right = sortedListToBST(quick.Next)
	return node
}
