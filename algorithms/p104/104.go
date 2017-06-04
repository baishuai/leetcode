package p104

/**
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes
 along the longest path from the root node down to the farthest leaf node.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	depth := 0
	var traversal func(l int, node *TreeNode)
	traversal = func(l int, node *TreeNode) {
		if node == nil {
			if l > depth {
				depth = l
			}
			return
		}
		l++
		traversal(l, node.Left)
		traversal(l, node.Right)
	}
	traversal(0, root)
	return depth
}
