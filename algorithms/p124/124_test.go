package p124

import "testing"

func test(t *testing.T, root *TreeNode, exp int) {
	if ans := maxPathSum(root); ans != exp {
		t.Fatal("error answer")
	}
}

func TestEXample0(t *testing.T) {
	test(t, &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 6)
}

func TestEXample1(t *testing.T) {
	test(t, &TreeNode{1, &TreeNode{-2, nil, nil}, &TreeNode{3, nil, nil}}, 4)
}

//func TestExample1(t *testing.T) {
//	test(t, []int{7, 6, 4, 3, 1}, 0)
//}
