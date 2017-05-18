package p144

import "testing"

func TestExample(t *testing.T) {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{Val: 3}, nil}}

	res := preorderTraversal(root)
	exp := []int{1, 2, 3}
	if len(res) != len(exp) {
		t.Fail()
	}
	for i, v := range exp {
		if res[i] != v {
			t.Error()
		}
	}
}
