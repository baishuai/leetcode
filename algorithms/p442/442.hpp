
#ifndef LEETCODE_442_HPP
#define LEETCODE_442_HPP

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
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array),
 some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
Input:
[4,3,2,7,8,2,3,1]

Output:
[2,3]
 */

class Solution {
public:
    vector<int> findDuplicates(vector<int> &nums) {
        vector<int> res;
        for (int i = 0; i < nums.size(); i++) {
            int n = abs(nums[i]) - 1;
            if (nums[n] < 0)
                res.emplace_back(abs(nums[i]));
            else
                nums[n] = -nums[n];
        }
        return res;
    }
};


#endif //LEETCODE_442_HPP
