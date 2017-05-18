class Solution(object):
    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        sum = 0
        for i in range(0,32):
            if n & 1 == 1:
                sum = sum + 1
            n = n >> 1
        return sum