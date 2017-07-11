//
// Created by bai on 17-6-25.
//

#ifndef LEETCODE_496_HPP
#define LEETCODE_496_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <unordered_map>
#include <stack>

using namespace std;


/**
  You are given two arrays (without duplicates) nums1 and nums2 where nums1’s
  elements are subset of nums2.
  Find all the next greater numbers for nums1's elements in the corresponding places of nums2.

The Next Greater Number of a number x in nums1 is
 the first greater number to its right in nums2. If it does not exist, output -1 for this number.
 */

class Solution {
public:
    vector<int> nextGreaterElement(vector<int> &findNums, vector<int> &nums) {
        stack<int> q;
        unordered_map<int, int> m;

        for (int n : nums) {
            while (q.size() > 0 && n > q.top()) {
                m[q.top()] = n;
                q.pop();
            }
            q.emplace(n);
        }

        vector<int> res;

        for (int n : findNums) {
            if (m.count(n) > 0) {
                res.emplace_back(m[n]);
            } else {
                res.emplace_back(-1);
            }
        }
        return res;
    }
};

#endif //LEETCODE_496_HPP
