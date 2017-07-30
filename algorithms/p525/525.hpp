
#ifndef LEETCODE_525_HPP
#define LEETCODE_525_HPP

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
#include <map>

using namespace std;

//Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.

class Solution {
public:
    int findMaxLength(vector<int> &nums) {
        map<int, int> s2pos;
        s2pos[0] = -1;
        int sum = 0;
        int max = 0;
        for (int i = 0; i < nums.size(); ++i) {
            sum += nums[i] == 0 ? -1 : 1;
            if (s2pos.count(sum) == 0)
                s2pos[sum] = i;
            else {
                max = std::max(max, i - s2pos[sum]);
            }
        }
        return max;
    }
};

#endif //LEETCODE_525_HPP
