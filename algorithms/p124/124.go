package p124

/**
Given a binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to
 any node in the tree along the parent-child connections.
 The path must contain at least one node and does not need to go through the root.

For example:
Given the below binary tree,

       1
      / \
     2   3
Return 6.
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

func maxPathSum(root *TreeNode) int {
	maxPath := 0
	if root != nil {
		maxPath = root.Val
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var maxTraversal func(node *TreeNode) int
	maxTraversal = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := maxTraversal(node.Left)
		right := maxTraversal(node.Right)

		singleHand := node.Val + max(left, right)
		if left+right+node.Val > maxPath {
			maxPath = left + right + node.Val
		}
		if singleHand > maxPath {
			maxPath = singleHand
		}
		if node.Val > maxPath {
			maxPath = node.Val
		}
		return max(singleHand, node.Val)
	}
	maxTraversal(root)
	return maxPath
}
