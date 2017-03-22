package p104

import "testing"

func TestExample(t *testing.T) {
	if maxDepth(nil) != 0 {
		t.Fail()
	}
}

func TestExample1(t *testing.T) {
	if maxDepth(&TreeNode{Val: 1}) != 1 {
		t.Fail()
	}
}
