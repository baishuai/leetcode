//
// Created by bai on 17-6-30.
//

#ifndef LEETCODE_238_HPP
#define LEETCODE_238_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <numeric>
#include <list>

using namespace std;

/**
Given an array of n integers where n > 1, nums,
 return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].

Solve it without division and in O(n).

For example, given [1,2,3,4], return [24,12,8,6].

Follow up:
Could you solve it with constant space complexity?
 (Note: The output array does not count as extra space for the purpose of space complexity analysis.)

 */

class Solution {
public:
    vector<int> productExceptSelf(vector<int> &nums) {

        vector<int> output(nums.size());
        for (int i = (int) (nums.size() - 1), acc = 1; i >= 0; --i) {
            output[i] = acc;
            acc *= nums[i];
        }

        for (int j = 0, acc = 1; j < output.size(); ++j) {
            output[j] *= acc;
            acc *= nums[j];
        }
        return output;
    }
};

#endif //LEETCODE_238_HPP
