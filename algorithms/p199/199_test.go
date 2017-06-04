package p199

import "testing"

func TestExample(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{2, nil, &TreeNode{Val: 5}},
		&TreeNode{3, nil, &TreeNode{Val: 4}}}

	res := rightSideView(root)
	exp := []int{1, 3, 4}
	if len(res) != len(exp) {
		t.Fatal("length", len(res))
	}
	for i, v := range exp {
		if res[i] != v {
			t.Error()
		}
	}
}
