package p099

/**
Two elements of a binary search tree (BST) are swapped by mistake.

Recover the tree without changing its structure.
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

func recoverTree(root *TreeNode) {
	// TODO Do not understand

	var first, second, prev *TreeNode
	var help func(node *TreeNode)
	help = func(node *TreeNode) {
		if node == nil {
			return
		}
		help(node.Left)
		if first == nil && prev != nil && prev.Val >= node.Val {
			first = prev
		}
		if first != nil && prev != nil && prev.Val > node.Val {
			second = node
		}
		prev = node
		help(node.Right)
	}

	help(root)
	first.Val, second.Val = second.Val, first.Val
}
