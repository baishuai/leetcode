package p092

import (
	"testing"
)

func test(t *testing.T, list *ListNode, m, n int, exp *ListNode) {
	ans := reverseBetween(list, m, n)
	for exp != nil {
		if ans == nil {
			t.Fatal("error answer length")
		}
		if exp.Val != ans.Val {
			t.Fatal("error answer squence")
		}
		exp = exp.Next
		ans = ans.Next
	}
	if ans != nil {
		t.Error("error length, longer")
	}
}

func TestExample0(t *testing.T) {
	list := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
						  &ListNode{5, nil}}}}}

	exp := &ListNode{1,
		&ListNode{4,
			&ListNode{3,
				&ListNode{2,
						  &ListNode{5, nil}}}}}

	test(t, list, 2, 4, exp)
}

func TestExample1(t *testing.T) {
	list := &ListNode{5, nil}
	exp := &ListNode{5, nil}

	test(t, list, 1, 1, exp)
}
