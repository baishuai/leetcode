package p147

/**
Sort a linked list using insertion sort.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	master := &ListNode{0, head}

	nextInsert := head.Next
	endInsert := head
	endInsert.Next = nil
	for nextInsert != nil {
		insert := nextInsert
		nextInsert = nextInsert.Next

		if insert.Val >= endInsert.Val {
			endInsert.Next = insert
			endInsert = endInsert.Next
			endInsert.Next = nil
		} else {
			cur := master
			for cur.Next != nil && cur.Next.Val < insert.Val {
				cur = cur.Next
			}
			if cur.Next != insert {
				insert.Next = cur.Next
				cur.Next = insert
			}
		}

	}

	return master.Next
}
