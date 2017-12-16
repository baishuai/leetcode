
#ifndef LEETCODE_540_HPP
#define LEETCODE_540_HPP

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
Given a sorted array consisting of only integers where
 every element appears twice except for one element which appears once.
 Find this single element that appears only once.

Example 1:
Input: [1,1,2,3,3,4,4,8,8]
Output: 2
Example 2:
Input: [3,3,7,7,10,11,11]
Output: 10
Note: Your solution should run in O(log n) time and O(1) space.
 */

class Solution {
public:
    int singleNonDuplicate(vector<int> &nums) {
        int i = 0, j = static_cast<int>(nums.size());
        while (i < j) {
            int m = i + (j - i) / 2;
            if (m > 0 && nums[m - 1] == nums[m])
                m = m - 1;
            else if (m < nums.size() - 1 && nums[m + 1] == nums[m]) {}
            else
                return nums[m];
            if (m % 2 == 0)
                i = m + 2;
            else
                j = m - 1;
        }
        return nums[i];
    }
};

#endif //LEETCODE_540_HPP
