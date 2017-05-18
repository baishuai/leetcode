package p111

/**
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along
 the shortest path from the root node down to the nearest leaf node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil {
		return 1 + minDepth(root.Right)
	} else if root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	return 1 + min(minDepth(root.Left), minDepth(root.Right))
}
