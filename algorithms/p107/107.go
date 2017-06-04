package p107

/**
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
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
	t, b := 0, len(res)-1
	for t < b {
		res[t], res[b] = res[b], res[t]
		t++
		b--
	}
	return res
}
