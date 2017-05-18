package p024

import "testing"

func TestExample(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	ans := swapPairs(list)
	index := 0
	for ans != nil {
		if index%2 == 0 && index+2 != ans.Val {
			t.Fatal("error answer")
		}
		if index%2 == 1 && index != ans.Val {
			t.Fatal("error answer")
		}
		index++
		ans = ans.Next
	}
}
