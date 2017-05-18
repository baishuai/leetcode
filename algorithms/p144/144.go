package p144

/**
Given a binary tree, return the preorder traversal of its nodes' values.

For example:
Given binary tree {1,#,2,3},
   1
    \
     2
    /
   3
return [1,2,3].

Note: Recursive solution is trivial, could you do it iteratively?
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

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	visit := func(v int) {
		res = append(res, v)
	}
	cur, prev := root, root

	for cur != nil {
		if cur.Left == nil {
			visit(cur.Val)
			cur = cur.Right
		} else {
			prev = cur.Left
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
			}

			if prev.Right == nil {
				visit(cur.Val)
				prev.Right = cur
				cur = cur.Left
			} else {
				cur = cur.Right
				prev.Right = nil
			}
		}
	}
	return res
}
