package p129

import "testing"

func TestExample(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil}}
	if sumNumbers(root) != 25 {
		t.Fail()
	}
}
