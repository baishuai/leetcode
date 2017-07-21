
#ifndef LEETCODE_643_HPP
#define LEETCODE_643_HPP

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
Given an array consisting of n integers,
 find the contiguous subarray of given length k that has the maximum average value.
 And you need to output the maximum average value.

Example 1:
Input: [1,12,-5,-6,50,3], k = 4
Output: 12.75
Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
Note:
1 <= k <= n <= 30,000.
Elements of the given array will be in the range [-10,000, 10,000].
 */

class Solution {
public:
    double findMaxAverage(vector<int> &nums, int k) {
        int sum = 0;
        for (int i = 0; i < k && i < nums.size(); ++i) {
            sum += nums[i];
        }

        double ave = static_cast<double >(sum) / k;
        for (int j = k; j < nums.size(); ++j) {
            sum += nums[j];
            sum -= nums[j - k];
            ave = max(ave, static_cast<double >(sum) / k);
        }
        return ave;
    }
};

#endif //LEETCODE_643_HPP
