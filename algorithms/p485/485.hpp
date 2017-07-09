//
// Created by bai on 17-6-27.
//

#ifndef LEETCODE_485_HPP
#define LEETCODE_485_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

using namespace std;

/**

 Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.
Note:

The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000

 */

class Solution {
public:
    int findMaxConsecutiveOnes(vector<int> &nums) {
        int max = 0, cur = 0;
        for (int n : nums) {
            if (n == 1) {
                cur++;
            } else if (cur > 0) {
                max = std::max(max, cur);
                cur = 0;
            }
        }
        max = std::max(max, cur);
        return max;
    }
};


#endif //LEETCODE_485_HPP
