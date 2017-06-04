package p145

/**
Given a binary tree, return the postorder traversal of its nodes' values.

For example:
Given binary tree {1,#,2,3},
   1
    \
     2
    /
   3
return [3,2,1].

Note: Recursive solution is trivial, could you do it iteratively?
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	root = &TreeNode{Left: root}
	visit := func(s, e *TreeNode) {
		vals := make([]int, 0)
		for s != e {
			vals = append(vals, s.Val)
			s = s.Right
		}
		vals = append(vals, e.Val)
		i, j := 0, len(vals)-1
		for i < j {
			vals[i], vals[j] = vals[j], vals[i]
			i++
			j--
		}
		res = append(res, vals...)
	}
	cur, prev := root, root

	for cur != nil {
		if cur.Left == nil {
			//visit(cur.Val)
			cur = cur.Right
		} else {
			prev = cur.Left
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
			}

			if prev.Right == nil {
				//visit(cur.Val)
				prev.Right = cur
				cur = cur.Left
			} else {
				visit(cur.Left, prev)
				cur = cur.Right
				prev.Right = nil
			}
		}
	}
	return res
}
