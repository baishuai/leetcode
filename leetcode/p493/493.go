package p493

/**
Given an array nums, we call (i, j) an important reverse pair if i < j and nums[i] > 2*nums[j].

You need to return the number of important reverse pairs in the given array.

Example1:

Input: [1,3,2,3,1]
Output: 2
Example2:

Input: [2,4,3,5,1]
Output: 3
Note:
The length of the given array will not exceed 50,000.
All the numbers in the input array are in the range of 32-bit integer.
 */

type bstNode struct {
	val   int
	sum   int
	left  *bstNode
	right *bstNode
}

func newNode(val int) *bstNode {
	return &bstNode{
		val: val,
		sum: 1,
	}
}

type bsTree struct {
	root *bstNode
}

func (t *bsTree) Search(v int) int {
	cnt := 0
	node := t.root
	for node != nil {
		if node.val < v {
			cnt += node.sum
			node = node.right
		} else {
			node = node.left
		}
	}
	return cnt
}

func (t *bsTree) Insert(v int) {
	if t.root == nil {
		t.root = newNode(v)
	} else {
		node := t.root
		for {
			if node.val < v {
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
				node.sum++
				break
			}
		}
	}
}

func reversePairs(nums []int) int {
	res := 0
	bst := bsTree{}
	for i := len(nums) - 1; i >= 0; i-- {
		res += bst.Search(nums[i])
		bst.Insert(nums[i] * 2)
	}
	return res
}
