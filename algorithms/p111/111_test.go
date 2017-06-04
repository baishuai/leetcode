package p111

import "testing"

func TestExample(t *testing.T) {
	if minDepth(nil) != 0 {
		t.Fail()
	}
}

func TestExample1(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			nil},
		nil}

	if minDepth(root) != 3 {
		t.Fail()
	}
}
