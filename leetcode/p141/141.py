"""
Leetcode p141, do not have golang environment
"""
# Given a linked list, determine if it has a cycle in it.

# Follow up:
# Can you solve it without using extra space?

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class ListNode(object):
    """
    Leetcode ListNode class
    """
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    """
    Leetcode sulution
    """
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        if head is None:
            return False
        slow = head
        quick = head.next
        while quick != None and quick != slow:
            slow = slow.next
            quick = quick.next
            if quick != None:
                quick = quick.next
        return quick != None
