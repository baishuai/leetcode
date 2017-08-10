
#ifndef LEETCODE_477_HPP
#define LEETCODE_477_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>
#include <stack>
#include <string>

using namespace std;

/*
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Now your job is to find the total Hamming distance between all pairs of the given numbers.

Example:
Input: 4, 14, 2

Output: 6

Explanation: In binary representation, the 4 is 0100, 14 is 1110, and 2 is 0010 (just
showing the four bits relevant in this case). So the answer will be:
HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.
Note:
Elements of the given array are in the range of 0 to 10^9
Length of the array will not exceed 10^4.
 */

class Solution {
public:
    int totalHammingDistance(vector<int> &nums) {
        vector<int> one(32), zero(32);
        for (int n : nums) {
            unsigned un = static_cast<unsigned int>(n);
            for (int i = 0; i < 32; ++i) {
                if ((un >> i) & 1) {
                    one[i]++;
                } else {
                    zero[i]++;
                }
            }
        }
        int total = 0;
        for (int j = 0; j < 32; ++j) {
            total += one[j] * zero[j];
        }
        return total;
    }
};

#endif //LEETCODE_477_HPP
