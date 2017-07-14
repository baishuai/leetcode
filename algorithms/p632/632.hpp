//
// Created by bai on 17-7-2.
//

#ifndef LEETCODE_632_HPP
#define LEETCODE_632_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>

using namespace std;

struct vinr {
    int value;
    int ridx;

    bool operator>(const vinr &rhs) const {
        return value > rhs.value;
    }

    vinr(int value, int ridx) : value(value), ridx(ridx) {}
};

class Solution {
public:
    vector<int> smallestRange(vector<vector<int>> &nums) {
        priority_queue<vinr, vector<vinr>, greater<vinr>> priq;

        int left = numeric_limits<int>::max(), right = numeric_limits<int>::min();
        for (int i = 0; i < nums.size(); ++i) {
            priq.push(vinr(nums[i].front(), i));
            if (nums[i].front() > right) {
                right = nums[i].front();
            }
            if (nums[i].front() < left) {
                left = nums[i].front();
            }
        }
        vector<int> idxs(nums.size(), 1);
        int minRange = right - left;
        int l = left, r = right;
        while (true) {
            auto idx = priq.top().ridx;
            priq.pop();
            if (idxs[idx] >= nums[idx].size()) {
                break;
            }
            int v = nums[idx][idxs[idx]];
            priq.push(vinr(v, idx));
            idxs[idx]++;

            if (v > right) {
                right = v;
            }
            left = priq.top().value;

            if (right - left < minRange) {
                l = left;
                r = right;
                minRange = r - l;
            }
        }
        return {l, r};
    }
};


#endif //LEETCODE_632_HPP
