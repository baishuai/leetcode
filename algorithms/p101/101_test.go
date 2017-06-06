package p101

import "testing"

func TestEample(t *testing.T) {
	isSymmetric(nil)
}

func Test1(t *testing.T) {
	isSymmetric(&TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 2}})
}

func Test2(t *testing.T) {
	isSymmetric(&TreeNode{1, &TreeNode{Val: 2}, nil})
}
