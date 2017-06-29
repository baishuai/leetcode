//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_283_HPP
#define LEETCODE_283_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    void moveZeroes(vector<int> &nums) {
        int pos = 0;
        for (int i = 0; i < nums.size(); ++i) {
            if (nums[i] != 0) {
                    nums[pos++] = nums[i];
            }
        }
        fill(nums.begin() + pos, nums.end(), 0);
    }
};

#endif //LEETCODE_283_HPP
