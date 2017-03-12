package p098

import "testing"

func test(t *testing.T, root *TreeNode, exp bool) {
	if isValidBST(root) != exp {
		t.Fatal("error answer")
	}
}

func TestExample(t *testing.T) {
	root := &TreeNode{
		2,
		&TreeNode{1, nil, nil},
		&TreeNode{3, nil, nil},
	}
	test(t, root, true)
}

func TestExtra0(t *testing.T) {
	root := &TreeNode{
		3,
		&TreeNode{1, nil, nil},
		&TreeNode{2, nil, nil},
	}
	test(t, root, false)
}

func TestExtra1(t *testing.T) {
	root := &TreeNode{
		3,
		&TreeNode{1, nil, nil},
		nil,
	}
	test(t, root, true)
}

func TestExtra2(t *testing.T) {
	root := &TreeNode{
		1,
		&TreeNode{1, nil, nil},
		nil,
	}
	test(t, root, false)
}

func TestExtra3(t *testing.T) {
	root := &TreeNode{
		10,
		&TreeNode{5, nil, nil},
		&TreeNode{15,
				  &TreeNode{6, nil, nil},
				  &TreeNode{20, nil, nil}},
	}
	test(t, root, false)
}
