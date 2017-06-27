//
// Created by bai on 17-6-26.
//

#ifndef LEETCODE_594_HPP
#define LEETCODE_594_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    int findLHS(vector<int> &nums) {
        unordered_map<int, int> nCnt;

        for (int n : nums) {
            nCnt[n]++;
        }
        if (nCnt.size() <= 1) {
            return 0;
        }
        int ans = 0;
        for_each(nCnt.begin(), nCnt.end(), [&ans, &nCnt](pair<const int &, int> p) {
            if (nCnt.count(p.first + 1) > 0) {
                ans = max(ans, p.second + nCnt[p.first + 1]);
            } else if (nCnt.count(p.first - 1) > 0) {
                ans = max(ans, p.second + nCnt[p.first - 1]);
            }
        });
        return ans;
    }
};


#endif //LEETCODE_594_HPP
