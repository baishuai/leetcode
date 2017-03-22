package p110

/**
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as a binary tree
 in which the depth of the two subtrees of every node never differ by more than 1.
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func heightAndBalance(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	lh, lb := heightAndBalance(node.Left)
	if lb {
		rh, rb := heightAndBalance(node.Right)
		if rb && abs(lh-rh) <= 1 {
			return max(lh, rh) + 1, true
		}
	}
	return 0, false
}

func isBalanced(root *TreeNode) bool {
	_, b := heightAndBalance(root)
	return b
}
