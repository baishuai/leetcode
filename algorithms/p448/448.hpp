//
// Created by bai on 17-6-27.
//

#ifndef LEETCODE_448_HPP
#define LEETCODE_448_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

using namespace std;

/**
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array),
 some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime?
 You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
 */

class Solution {
public:
    vector<int> findDisappearedNumbers(vector<int> &nums) {
        for (int n: nums) {
            n = abs(n) - 1;
            if (nums[n] > 0) {
                nums[n] = 0 - nums[n];
            }
        }
        vector<int> res;
        for (int i = 0; i < nums.size(); ++i) {
            if (nums[i] > 0) {
                res.emplace_back(i + 1);
            }
        }
        return res;
    }
};


#endif //LEETCODE_448_HPP
