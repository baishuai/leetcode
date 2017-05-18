# Definition for binary tree with next pointer.
class TreeLinkNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
        self.next = None

class Solution:
    # @param root, a tree link node
    # @return nothing
    def connect(self, root):
        if root is None:
            return
        pre = root
        cur = None
        while pre.left is not None:
            cur = pre
            while cur is not None:
                cur.left.next = cur.right
                if cur.next is not None:
                    cur.right.next = cur.next.left
                cur = cur.next
            pre = pre.left


        