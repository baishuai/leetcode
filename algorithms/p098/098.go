package p098

/**
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
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

func isValidBST(root *TreeNode) bool {
	return isValidPart(root, nil, nil)
}

func isValidPart(node, min, max *TreeNode) bool {
	if node == nil {
		return true
	}
	if (min != nil && node.Val <= min.Val) || (max != nil && node.Val >= max.Val) {
		return false
	}
	return isValidPart(node.Left, min, node) && isValidPart(node.Right, node, max)
}
