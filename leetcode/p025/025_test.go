package p025

import (
	"fmt"
	"testing"
)

func Test2(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	exp := &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}}
	ans := reverseKGroup(list, 2)
	for exp != nil {
		if ans == nil {
			t.Fatal("error answer length")
		}
		if exp.Val != ans.Val {
			t.Fatal("error answer squence")
		}
		fmt.Println(ans.Val)
		exp = exp.Next
		ans = ans.Next
	}
}
