package p129

/**
Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path 1->2->3 which represents the number 123.

Find the total sum of all root-to-leaf numbers.

For example,

    1
   / \
  2   3
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.

Return the sum = 12 + 13 = 25.
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

func sumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(node *TreeNode, num int)
	dfs = func(node *TreeNode, num int) {
		if node != nil {
			num = num*10 + node.Val
			dfs(node.Left, num)
			dfs(node.Right, num)
			if node.Left == nil && node.Right == nil {
				sum += num
			}
		}

	}
	dfs(root, 0)
	return sum
}
