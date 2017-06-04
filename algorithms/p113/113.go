package p113

/**
Given a binary tree and a sum, find all root-to-leaf paths where
 each path's sum equals the given sum.

For example:
Given the below binary tree and sum = 22,
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
return
[
   [5,4,11,2],
   [5,8,4,5]
]
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

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(node *TreeNode, s int)
	dfs = func(node *TreeNode, s int) {
		if node == nil {
			return
		} else if node.Left == nil && node.Right == nil {
			if s+node.Val == sum {
				pathOk := make([]int, len(path)+1)
				copy(pathOk, path)
				pathOk[len(path)] = node.Val
				res = append(res, pathOk)
			}
		} else {
			s += node.Val
			if s > sum {
				return
			}
			path = append(path, node.Val)
			dfs(node.Left, s)
			dfs(node.Right, s)
			path = path[:len(path)-1]
		}
	}
	dfs(root, 0)
	return res
}
