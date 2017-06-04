package p606

import (
	"bytes"
	"strconv"
)

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

func tree2str(t *TreeNode) string {

	res := new(bytes.Buffer)

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res.WriteString(strconv.FormatInt(int64(node.Val), 10))
		if node.Left != nil {
			res.WriteByte('(')
			preorder(node.Left)
			res.WriteByte(')')
		}
		if node.Right != nil {
			if node.Left == nil {
				res.WriteString("()")
			}
			res.WriteByte('(')
			preorder(node.Right)
			res.WriteByte(')')
		}
	}
	preorder(t)
	return res.String()
}
