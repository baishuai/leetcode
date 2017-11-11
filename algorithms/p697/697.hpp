
#ifndef LEETCODE_697_HPP
#define LEETCODE_697_HPP

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

class CntPos {
public:
    int cnt;
    int left;
    int right;

    void update(int pos) {
        if (this->cnt == 0) {
            this->left = pos;
        }
        this->right = pos;
        cnt++;
    }
};

class Solution {
public:
    int findShortestSubArray(vector<int> &nums) {
        unordered_map<int, CntPos> cnts;
        for (int i = 0; i < nums.size(); ++i) {
            cnts[nums[i]].update(i);
        }
        int mcnt = 1, mlen = 1;
        for (auto &p : cnts) {
            if (p.second.cnt > mcnt) {
                mcnt = p.second.cnt;
                mlen = p.second.right - p.second.left + 1;
            } else if (p.second.cnt == mcnt && p.second.right - p.second.left + 1 < mlen) {
                mcnt = p.second.cnt;
                mlen = p.second.right - p.second.left + 1;
            }
        }
        return mlen;
    }
};

#endif //LEETCODE_697_HPP
