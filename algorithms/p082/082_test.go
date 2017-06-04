package p082

import (
	"testing"
)

func test(t *testing.T, head, exp *ListNode) {
	ans := deleteDuplicates(head)

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
	test(t, &ListNode{1,
		&ListNode{1,
			&ListNode{1,
				&ListNode{2,
					&ListNode{3, nil}}}}},
		&ListNode{2, &ListNode{3, nil}})
}

func TestExtra0(t *testing.T) {
	test(t, &ListNode{1,
		&ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{4,
							&ListNode{5,
								&ListNode{6,
									&ListNode{7,
										&ListNode{7, nil}}}}}}}}}},
		&ListNode{2, &ListNode{3, &ListNode{5, &ListNode{6, nil}}}})
}

func TestExtra1(t *testing.T) {
	test(t, &ListNode{1, &ListNode{1, nil}}, nil)
}
