package p106

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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	node := &TreeNode{Val: val}
	for i, v := range inorder {
		if v == val {
			node.Left = buildTree(inorder[:i], postorder[:i])
			node.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
			break
		}
	}
	return node
}
