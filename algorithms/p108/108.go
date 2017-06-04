package p108

/**
Given an array where elements are sorted in ascending order,
 convert it to a height balanced BST.

Subscribe to see which companies asked this question.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	nlen := len(nums)
	if nlen == 0 {
		return nil
	} else if nlen == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	mid := nlen / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])
	return node
}
