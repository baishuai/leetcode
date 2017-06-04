# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):


def isBadVersion(version):
    return version > 5


class Solution(object):
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        i = 0
        j = n
        while i < j:
            h = i + (j - i) / 2
            if not isBadVersion(h):
                i = h + 1
            else:
                j = h

        return i
