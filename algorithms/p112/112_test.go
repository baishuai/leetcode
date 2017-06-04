package p112

import "testing"

func TestExample(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			nil},
		nil}

	if hasPathSum(root, 6) != true {
		t.Fail()
	}

	if hasPathSum(root, 3) == true {
		t.Fail()
	}
}
