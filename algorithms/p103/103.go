package p103

/**
Given a binary tree, return the zigzag level order traversal of its nodes' values.
 (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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

	for i := 1; i < len(res); i += 2 {
		reverse := res[i]
		l, h := 0, len(reverse)-1
		for l < h {
			reverse[l], reverse[h] = reverse[h], reverse[l]
			l++
			h--
		}
	}
	return res
}
