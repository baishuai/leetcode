package p147

import "testing"

func test(t *testing.T, head *ListNode, size int) {
	head = insertionSortList(head)
	for head != nil && head.Next != nil {
		if head.Val > head.Next.Val {
			t.Fatal("error value")
		}
		head = head.Next
		size--
	}
	if size != 1 {
		t.Fatal("size error")
	}
}

func TestExample0(t *testing.T) {
	test(t, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 4)
}

func TestExample1(t *testing.T) {
	test(t, &ListNode{4, &ListNode{3, &ListNode{3, &ListNode{4, nil}}}}, 4)
}

func TestExample2(t *testing.T) {
	test(t, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}, 4)
}
