package p199

/**
Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

For example:
Given the following binary tree,
   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
You should return [1, 3, 4].


 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)

	visit := func(v, h int) {
		if h == len(res) {
			res = append(res, v)
		} else {
			res[h] = v
		}
	}
	cur, prev := root, root
	height := 0
	for cur != nil {
		if cur.Left == nil {
			visit(cur.Val, height)
			cur = cur.Right
			height++
		} else {
			prev = cur.Left
			count := 1
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
				count ++
			}

			if prev.Right == nil {
				visit(cur.Val, height)
				prev.Right = cur
				cur = cur.Left
				height++
			} else {
				height -= count
				cur = cur.Right
				prev.Right = nil
			}
		}
	}
	return res
}
