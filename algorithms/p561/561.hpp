//
// Created by bai on 17-6-26.
//

#ifndef LEETCODE_561_HPP
#define LEETCODE_561_HPP


#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

using namespace std;

/**
Given an array of 2n integers, your task is to group these integers into n pairs of integer,
 say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

Example 1:
Input: [1,4,3,2]

Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).
Note:
n is a positive integer, which is in the range of [1, 10000].
All the integers in the array will be in the range of [-10000, 10000].
 */

class Solution {
public:
    int arrayPairSum(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        int sum = 0;
        for (int i = 0; i < nums.size(); ++i) {
            sum += nums[i++];
        }
        return sum;
    }
};

#endif //LEETCODE_561_HPP
