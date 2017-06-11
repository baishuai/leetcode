package p617

/**
Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

Example 1:
Input:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
Output:
Merged tree:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
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

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	res := &TreeNode{}
	mergeNode(res, t1, t2)
	return res
}

func mergeNode(dst, n1, n2 *TreeNode) {
	dst.Val = 0
	if n1 != nil {
		dst.Val += n1.Val
	}
	if n2 != nil {
		dst.Val += n2.Val
	}

	var l1, l2 *TreeNode
	var r1, r2 *TreeNode
	if n1 != nil {
		l1 = n1.Left
		r1 = n1.Right
	}
	if n2 != nil {
		l2 = n2.Left
		r2 = n2.Right
	}

	if l1 != nil || l2 != nil {
		dst.Left = new(TreeNode)
		mergeNode(dst.Left, l1, l2)
	}
	if r1 != nil || r2 != nil {
		dst.Right = new(TreeNode)
		mergeNode(dst.Right, r1, r2)
	}
}
