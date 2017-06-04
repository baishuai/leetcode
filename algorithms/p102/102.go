package p102

/**
Given a binary tree, return the level order traversal of its
 nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
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

func levelOrder(root *TreeNode) [][]int {

	res := make([][]int, 0)

	var traversal func(l int, node *TreeNode)
	traversal = func(l int, node *TreeNode) {
		if node == nil {
			return
		}
		if len(res) <= l {
			res = append(res, []int{node.Val})
		} else {
			res[l] = append(res[l], node.Val)
		}
		traversal(l+1, node.Left)
		traversal(l+1, node.Right)
	}
	traversal(0, root)
	return res
}
