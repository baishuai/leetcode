//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_565_HPP
#define LEETCODE_565_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    int arrayNesting(vector<int> &nums) {
        int maxSize = 0;
        for (int i = 0; i < nums.size(); ++i) {
            if (nums[i] < 0) {
                continue;
            }
            int p = nums[i];
            int q = p;
            int cnt = 1;
            while (p != nums[q]) {
                cnt++;
                int t = q;
                q = nums[q];
                nums[t] = -1;
            }
            maxSize = max(maxSize, cnt);
        }
        return maxSize;
    }
};

#endif //LEETCODE_565_HPP
