package p623

import "fmt"

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

func dfs(node *TreeNode, v int, d int) {
	if node == nil {
		return
	}
	if d > 1 {
		dfs(node.Left, v, d-1)
		dfs(node.Right, v, d-1)
		return
	}
	fmt.Println(node.Val)
	ol, or := node.Left, node.Right

	node.Left = &TreeNode{Val: v}
	node.Right = &TreeNode{Val: v}
	node.Left.Left = ol
	node.Right.Right = or
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		node := &TreeNode{Val: v}
		node.Left = root
		return node
	}

	dfs(root, v, d-1)
	return root

}
