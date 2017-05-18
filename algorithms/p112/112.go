package p112

/**
Given a binary tree and a sum, determine if the tree has a root-to-leaf path
 such that adding up all the values along the path equals the given sum.

For example:
Given the below binary tree and sum = 22,
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
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

func hasPathSum(root *TreeNode, sum int) bool {
	has := false

	var dfs func(node *TreeNode, s int)
	dfs = func(node *TreeNode, s int) {
		if !has {
			if node == nil {
				return
			} else if node.Left == nil && node.Right == nil {
				if s+node.Val == sum {
					has = true
				}
			} else {
				s += node.Val
				dfs(node.Left, s)
				dfs(node.Right, s)
			}
		}
	}
	dfs(root, 0)
	return has
}
