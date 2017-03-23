package p113

import "testing"

func TestExample(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{2,
				  &TreeNode{3, nil, nil},
				  nil},
		nil}

	if len(pathSum(root, 6)) != 1 {
		t.Fail()
	}
}
