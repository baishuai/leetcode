//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_560_HPP
#define LEETCODE_560_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <set>
#include <unordered_set>
#include <unordered_map>

using namespace std;

/**
 *
Given an array of integers and an integer k,
 you need to find the total number of continuous subarrays whose sum equals to k.

Example 1:
Input:nums = [1,1,1], k = 2
Output: 2
Note:
The length of the array is in range [1, 20,000].
The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].

 */

class Solution {
public:
    int subarraySum(vector<int> &nums, int k) {
        unordered_map<int, int> sums;
        sums[0]++;
        int acc = 0, cnt = 0;
        for (int i = 0; i < nums.size(); ++i) {
            acc += nums[i];
            if (sums[acc - k] > 0) {
                cnt += sums[acc - k];
            }
            sums[acc]++;
        }
        return cnt;
    }
};


#endif //LEETCODE_560_HPP
