package p352

/**
Given a data stream input of non-negative integers a1, a2, ..., an, ...,
 summarize the numbers seen so far as a list of disjoint intervals.

For example, suppose the integers from the data stream are 1, 3, 7, 2, 6, ...,
 then the summary will be:

[1, 1]
[1, 1], [3, 3]
[1, 1], [3, 3], [7, 7]
[1, 3], [7, 7]
[1, 3], [6, 7]
Follow up:
What if there are lots of merges and the number of disjoint intervals are small
 compared to the data stream's size?
 */

/**
 * Definition for an interval.
**/
type Interval struct {
	Start int
	End   int
}

type BSTreeNode struct {
	val   Interval
	left  *BSTreeNode
	right *BSTreeNode
}

func newNode(v int) *BSTreeNode {
	node := BSTreeNode{
		val: Interval{v, v},
	}
	return &node
}

func findMin(node *BSTreeNode) *BSTreeNode {
	if node == nil {
		return nil
	} else if node.left == nil {
		return node
	} else {
		return findMin(node.left)
	}
}

func addKey(val int, node *BSTreeNode) *BSTreeNode {

	if node == nil {
		node = newNode(val)
	} else if node.val.Start > val {
		node.left = addKey(val, node.left)
	} else if node.val.End < val {
		node.right = addKey(val, node.right)
	}
	return node
}

func remove(x *Interval, node *BSTreeNode) *BSTreeNode {
	if node == nil || x == nil {
		return node
	} else if x.Start > node.val.End {
		node.right = remove(x, node.right)
	} else if x.End < node.val.Start {
		node.left = remove(x, node.left)
	} else if node.left != nil && node.right != nil {
		node.val = findMin(node.right).val
		node.right = remove(&node.val, node.right)
	} else {
		if node.left != nil {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

type SummaryRanges struct {
	size int
	root *BSTreeNode
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{
		size: 0,
		root: nil,
	}
}

func findKey(val int, node *BSTreeNode) *BSTreeNode {
	if node == nil {
		return nil
	}
	if node.val.Start > val {
		return findKey(val, node.left)
	} else if node.val.End < val {
		return findKey(val, node.right)
	} else {
		return node
	}
}

func (this *SummaryRanges) Addnum(val int) {
	if this.root == nil {
		this.size++
		this.root = newNode(val)
	} else {
		if findKey(val, this.root) != nil {
			return
		}
		left := findKey(val-1, this.root)
		right := findKey(val+1, this.root)
		if left == nil && right == nil {
			this.root = addKey(val, this.root)
			this.size++
		} else if left != nil && right == nil {
			left.val.End++
		} else if left == nil && right != nil {
			right.val.Start--
		} else {
			end := right.val.End
			this.root = remove(&right.val, this.root)
			left.val.End = end
			this.size--
		}
	}
}

func (this *SummaryRanges) Getintervals() []Interval {
	ans := make([]Interval, 0, this.size)

	var inorder func(node *BSTreeNode)
	inorder = func(node *BSTreeNode) {
		if node != nil {
			inorder(node.left)
			ans = append(ans, node.val)
			inorder(node.right)
		}
	}

	inorder(this.root)
	return ans
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Addnum(val);
 * param_2 := obj.Getintervals();
 */
