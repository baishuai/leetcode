package p061

import "testing"

func test(t *testing.T, head *ListNode, k int, exp *ListNode) {
	ans := rotateRight(head, k)
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
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	exp := &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}
	test(t, head, 2, exp)
}

func TestExtra1(t *testing.T) {
	head := &ListNode{1, nil}
	exp := &ListNode{1, nil}
	test(t, head, 2, exp)
}

func TestExtra2(t *testing.T) {
	head := &ListNode{1, &ListNode{2, nil}}
	exp := &ListNode{2, &ListNode{1, nil}}
	test(t, head, 1, exp)
}

func TestExtra3(t *testing.T) {
	head := &ListNode{1, nil}
	exp := &ListNode{1, nil}
	test(t, head, 0, exp)
}

func TestExtra4(t *testing.T) {
	head := &ListNode{1, nil}
	exp := &ListNode{1, nil}
	test(t, head, 1, exp)
}

func TestExtra5(t *testing.T) {
	var head *ListNode
	var exp *ListNode
	test(t, head, 0, exp)
}


func TestExtra6(t *testing.T) {
	head := &ListNode{1, nil}
	exp := &ListNode{1, nil}
	test(t, head, 99, exp)
}
