# package p138


# A linked list is given such that each node contains an additional random pointer which could point to any node in the list or null.

# Return a deep copy of the list

# Definition for singly-linked list with a random pointer.
class RandomListNode(object):
    def __init__(self, x):
        self.label = x
        self.next = None
        self.random = None

class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: RandomListNode
        :rtype: RandomListNode
        """
        cnt = 0
        cur = head
        while cur:
            cnt += 1
            cur = cur.next
        lst = [RandomListNode(i) for i in range(0, cnt)]
        cur = head
        for i in range(0, cnt):
            lst[i].label = cur.label
            cur.label = i
            if i+1 < cnt:
                lst[i].next = lst[i+1]
            cur = cur.next
        cur = head
        for i in range(0, cnt):
            if cur.random is not None:
                lst[i].random = lst[cur.random.label]
            cur = cur.next
        cur = head
        for i in range(0, cnt):
            cur.label = lst[i].label
            cur = cur.next

        return None if cnt == 0 else lst[0]


if __name__ == "__main__":
    solution = Solution()
    solution.copyRandomList(RandomListNode(-1))
