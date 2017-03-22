package p110

import "testing"

func TestExample(t *testing.T) {
	if isBalanced(nil) != true {
		t.Fail()
	}
}

func TestExample1(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{2,
				  &TreeNode{3, nil, nil},
				  nil},
		nil}

	if isBalanced(root) == true {
		t.Fail()
	}
}
