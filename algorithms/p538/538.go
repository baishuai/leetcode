package p538

/**
Given a Binary Search Tree (BST), convert it to a Greater Tree such that
 every key of the original BST is changed to the original key
 plus sum of all keys greater than the original key in BST.

Example:

Input: The root of a Binary Search Tree like this:
              5
            /   \
           2     13

Output: The root of a Greater Tree like this:
             18
            /   \
          20     13
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

func dfsConvert(node *TreeNode, sum int) int {
	if node == nil {
		return sum
	}
	node.Val += dfsConvert(node.Right, sum)
	return dfsConvert(node.Left, node.Val)
}

func convertBST(root *TreeNode) *TreeNode {
	dfsConvert(root, 0)
	return root
}
