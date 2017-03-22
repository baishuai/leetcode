package p105

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	node := &TreeNode{Val: val}
	for i, v := range inorder {
		if v == val {
			node.Left = buildTree(preorder[1:1+i], inorder[:i])
			node.Right = buildTree(preorder[1+i:], inorder[1+i:])
			break
		}
	}
	return node
}
