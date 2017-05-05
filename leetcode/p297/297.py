# Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

# Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

# For example, you may serialize the following tree

#     1
#    / \
#   2   3
#      / \
#     4   5

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        :type root: TreeNode
        :rtype: str
        """
        bfsQ = []
        seri = []
        if root is not None:
            bfsQ.append(root)
            seri.append(str(root.val))
        else:
            seri.append('null')
        while bfsQ:
            head = bfsQ[0]
            bfsQ = bfsQ[1:]
            if head.left is not None:
                bfsQ.append(head.left)
                seri.append(str(head.left.val))
            else:
                seri.append('null')
            if head.right is not None:
                bfsQ.append(head.right)
                seri.append(str(head.right.val))
            else:
                seri.append('null')
        result = '[' + ','.join(seri) +']'
        print(result)
        return '[' + ','.join(seri) +']'
        

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        :type data: str
        :rtype: TreeNode
        """
        data = data.lstrip('[').rstrip(']')
        nodes = data.split(',')
        print(nodes)
        root = None
        if len(nodes) == 0 or nodes[0] == 'null':
            return root
        else:
            root = TreeNode(int(nodes[0]))
            nodes = nodes[1:]
        bfsQ = [root]
        itx = 0
        while itx < len(nodes):
            node = nodes[itx]
            head = bfsQ[0]
            bfsQ = bfsQ[1:]
            if node == 'null':
                head.left = None
            else:
                head.left = TreeNode(int(node))
                bfsQ.append(head.left)
            itx += 1
            node = nodes[itx]
            if node == 'null':
                head.right = None
            else:
                head.right = TreeNode(int(node))
                bfsQ.append(head.right)
            itx += 1
        return root
# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))


if __name__ == "__main__":
    codec = Codec()
    root = None
    codec.deserialize(codec.serialize(root))