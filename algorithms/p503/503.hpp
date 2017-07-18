
#ifndef LEETCODE_503_HPP
#define LEETCODE_503_HPP

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
Given a circular array (the next element of the last element is the first element of the array),
 print the Next Greater Number for every element.
 The Next Greater Number of a number x is the first greater number to its traversing-order next in the array,
 which means you could search circularly to find its next greater number.
 If it doesn't exist, output -1 for this number.

Example 1:
Input: [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2;
The number 2 can't find next greater number;
The second 1's next greater number needs to search circularly, which is also 2.
Note: The length of given array won't exceed 10000.
 */

class Solution {
public:
    vector<int> nextGreaterElements(vector<int> &nums) {
        stack<int> q;
        vector<int> res(nums.size(), -1);
        auto len = nums.size();
        for (int i = 0; i < len * 2; i++) {
            int n = nums[i % len];
            while (q.size() > 0 && n > nums[q.top()]) {
                res[q.top()] = n;
                q.pop();
            }
            if (i < len)
                q.emplace(i);
        }
        return res;
    }
};

#endif //LEETCODE_503_HPP
