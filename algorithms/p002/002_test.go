package p002

import "testing"

func TestAddTwoNumbers(t *testing.T) {

	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	ok := []int{7, 0, 8}
	index := 0
	for result != nil {
		if result.Val != ok[index] {
			t.Fatal("error result", result.Val, ok[index])
		}
		result = result.Next
		index++
	}
}
