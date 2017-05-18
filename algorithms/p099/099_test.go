package p099

import (
	"testing"
)

func cmpTree(a, b *TreeNode) bool {

	if (a == nil) && ( b == nil) {
		return true
	} else if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return cmpTree(a.Left, b.Left) && cmpTree(a.Right, b.Right)
}

func test(t *testing.T, errbst, exp *TreeNode) {
	recoverTree(errbst)
	if !cmpTree(errbst, exp) {
		t.Error("error recover")
	}
}

func TestExample(t *testing.T) {
	errbst := &TreeNode{0,
						&TreeNode{1, nil, nil},
						nil}
	exp := &TreeNode{1,
					 &TreeNode{0, nil, nil},
					 nil}
	test(t, errbst, exp)
}

func TestExtra0(t *testing.T) {
	errbst :=  &TreeNode{
		10,
		&TreeNode{14, nil, nil},
		&TreeNode{15,
				  &TreeNode{6, nil, nil},
				  &TreeNode{20, nil, nil}},
	}
	exp := &TreeNode{
		10,
		&TreeNode{6, nil, nil},
		&TreeNode{15,
				  &TreeNode{14, nil, nil},
				  &TreeNode{20, nil, nil}},
	}

	test(t, errbst, exp)
}
