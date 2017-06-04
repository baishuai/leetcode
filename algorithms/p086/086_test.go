package p086

import "testing"

func test(t *testing.T, head *ListNode, x int, exp *ListNode) {
	ans := partition(head, x)
	for exp != nil {
		if ans == nil {
			t.Fatal("error answer length, shorter than exp")
		}
		if ans.Val != exp.Val {
			t.Fatal("error answer val")
		}
		exp = exp.Next
		ans = ans.Next
	}
	if ans != nil {
		t.Fatal("error answer length, longer than exp")
	}

}

func TestExample(t *testing.T) {
	head := &ListNode{1,
		&ListNode{4,
			&ListNode{3,
				&ListNode{2,
					&ListNode{5,
						&ListNode{2, nil}}}}}}
	exp := &ListNode{1,
		&ListNode{2,
			&ListNode{2,
				&ListNode{4,
					&ListNode{3,
						&ListNode{5, nil}}}}}}

	test(t, head, 3, exp)
}
