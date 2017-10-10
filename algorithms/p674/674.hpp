
#ifndef LEETCODE_674_HPP
#define LEETCODE_674_HPP

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

class Solution {
public:
    int findLengthOfLCIS(vector<int> &nums) {
        if (nums.empty())
            return 0;
        int max_len = 0;
        int prev = nums.front();
        int len = 1;
        for (int n : nums) {
            if (n > prev)
                len++;
            else {
                max_len = max(max_len, len);
                len = 1;
            }
            prev = n;
        }
        max_len = max(max_len, len);
        return max_len;
    }
};

#endif //LEETCODE_674_HPP
