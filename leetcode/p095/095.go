package p095

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

func generateTrees(n int) []*TreeNode {
	if n==0 {
		return nil
	}
	return genTrees(1, n)
}

func genTrees(s, e int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if s > e {
		res = append(res, nil)
	}
	for i := s; i <= e; i++ {
		leftSubTrees := genTrees(s, i-1)
		rightSubTrees := genTrees(i+1, e)
		for _, vl := range leftSubTrees {
			for _, vr := range rightSubTrees {
				root := &TreeNode{i, vl, vr}
				res = append(res, root)
			}
		}
	}
	return res
}
