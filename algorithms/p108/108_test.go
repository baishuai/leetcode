package p108

import "testing"

func printTree(t *testing.T, node *TreeNode) {
	if node == nil {
		return
	}
	printTree(t, node.Left)
	printTree(t, node.Right)
}

func TestExample(t *testing.T) {
	root := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7})
	printTree(t, root)
}

func TestExample1(t *testing.T) {
	root := sortedArrayToBST([]int{1, 2})
	printTree(t, root)
}
