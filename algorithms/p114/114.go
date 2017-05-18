package p114

/**
Given a binary tree, flatten it to a linked list in-place.

For example,
Given

         1
        / \
       2   5
      / \   \
     3   4   6
The flattened tree should look like:
   1
    \
     2
      \
       3
        \
         4
          \
           5
            \
             6

 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
rotate node.left to node.right
link pre node.right to pre node's pre node
 */
func rotate(node *TreeNode) {
	// node.left must not be nil
	left := node.Left
	pre := left
	for pre.Right != nil {
		pre = pre.Right
	}
	pre.Right = node.Right
	node.Left = nil
	node.Right = left
}

func flatten(root *TreeNode) {
	//rotate
	cur := root
	for cur != nil {
		if cur.Left != nil {
			rotate(cur)
		}
		cur = cur.Right
	}
}
