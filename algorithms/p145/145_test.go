package p145

import "testing"

func TestExample(t *testing.T) {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{Val: 3}, nil}}

	res := postorderTraversal(root)
	exp := []int{3, 2, 1}
	if len(res) != len(exp) {
		t.Fatal(len(res), len(exp))
	}
	for i, v := range exp {
		if res[i] != v {
			t.Error()
		}
	}
}
