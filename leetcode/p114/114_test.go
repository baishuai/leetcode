package p114

import "testing"

func TestExample(t *testing.T) {
	root := &TreeNode{1,
					  &TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 4}},
					  &TreeNode{5, nil, &TreeNode{Val: 6}}}

	flatten(root)

	for i := 1; i <= 6; i++ {
		if root.Val != i || root.Left != nil {
			t.Fail()
		}
		root = root.Right
	}
}
