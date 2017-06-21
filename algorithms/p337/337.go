package p337

/**
The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

Example 1:
     3
    / \
   2   3
    \   \
     3   1
Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
Example 2:
     3
    / \
   4   5
  / \   \
 1   3   1
Maximum amount of money the thief can rob = 4 + 5 = 9.

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

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	l1, l2 := dfs(node.Left)
	r1, r2 := dfs(node.Right)
	return node.Val + l2 + r2, max(l1, l2) + max(r1, r2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(root *TreeNode) int {
	n1, n2 := dfs(root)
	return max(n1, n2)
}
