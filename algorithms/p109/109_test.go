package p109

import "testing"

func TestExamplo0(t *testing.T) {
	sortedListToBST(&ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5, nil}}}}})
}
