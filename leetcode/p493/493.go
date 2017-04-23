package p493

import (
	"sort"
)

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

func mergeSort(nums, helper []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	cnt := mergeSort(nums, helper, start, mid) + mergeSort(nums, helper, mid+1, end)

	copy(helper, nums[start:mid+1])
	ix, ls, rs := start, start, mid+1
	offset := 0
	for ix <= end {
		if ls > mid || (rs <= end && helper[ls-start] >= nums[rs]) {
			if start+offset <= mid {
				find := sort.SearchInts(helper[offset:mid+1-start], nums[rs]*2+1)
				cnt += mid + 1 - (start + offset) - find
				offset += find
			}
			nums[ix] = nums[rs]
			ix++
			rs++
		} else {
			nums[ix] = helper[ls-start]
			ix++
			ls++
		}
	}
	return cnt
}

func reversePairs(nums []int) int {
	helper := make([]int, len(nums)/2+1)
	return mergeSort(nums, helper, 0, len(nums)-1)
}
