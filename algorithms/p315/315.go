package p315

/**
You are given an integer array nums and you have to return a new counts array. The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].

Example:

Given nums = [5, 2, 6, 1]

To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.
Return the array [2, 1, 1, 0].
 */

//BST
type bstNode struct {
	val   int
	dup   int
	sum   int
	left  *bstNode
	right *bstNode
}

func newNode(val int) *bstNode {
	return &bstNode{
		val: val,
		dup: 1,
	}
}

type bsTree struct {
	root *bstNode
}

func (t *bsTree) Insert(v int) int {
	cnt := 0
	if t.root == nil {
		t.root = newNode(v)
	} else {
		node := t.root
		for {
			if node.val < v {
				cnt += (node.sum + node.dup)
				if node.right == nil {
					node.right = newNode(v)
					break
				} else {
					node = node.right
				}
			} else if node.val > v {
				node.sum++
				if node.left == nil {
					node.left = newNode(v)
					break
				} else {
					node = node.left
				}
			} else {
				node.dup++
				cnt += node.sum
				break
			}
		}
	}
	return cnt
}

func countSmaller(nums []int) []int {

	res := make([]int, len(nums))
	bst := bsTree{}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = bst.Insert(nums[i])
	}
	return res
}
