
# Reverse bits of a given 32 bits unsigned integer.

# For example, given input 43261596 (represented in binary as 00000010100101000001111010011100),
# return 964176192 (represented in binary as 00111001011110000010100101000000).

# Follow up:
# If this function is called many times, how would you optimize it?

# Related problem: Reverse Integer

class Solution:
    # @param n, an integer
    # @return an integer
    def reverseBits(self, n):
        if n == 0:
            return 0
        result = 0
        for _ in range(0, 32):
            result <<= 1
            if (n & 1) == 1:
                result += 1
            n >>= 1
        return result

if __name__ == "__main__":
    s = Solution()
    print( s.reverseBits(43261596))