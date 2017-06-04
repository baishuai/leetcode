package p101

/**
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
But the following [1,2,2,null,3,null,3] is not:
    1
   / \
  2   2
   \   \
   3    3
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricLR(root.Left, root.Right)
}

func isSymmetricLR(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil) != (right == nil) {
		return false
	}
	return (left.Val == right.Val) &&
		isSymmetricLR(left.Left, right.Right) &&
		isSymmetricLR(left.Right, right.Left)
}
