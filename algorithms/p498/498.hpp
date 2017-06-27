//
// Created by bai on 17-6-27.
//

#ifndef LEETCODE_498_HPP
#define LEETCODE_498_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    vector<int> findDiagonalOrder(vector<vector<int>> &matrix) {
        auto m = matrix.size();
        if (m == 0) {
            return vector<int>();
        }
        auto n = matrix[0].size();

        vector<int> res;
        res.resize(m * n);
        int id = 0;
        for (int i = 0; i < m + n - 1; ++i) {
            int lb = max(0, (int) (i - n + 1));
            int ub = min(i, (int) m - 1);

            if (i % 2 == 0) {
                for (int j = ub; j >= lb; --j) {
                    res[id++] = matrix[j][i - j];
                }
            } else {
                for (int j = lb; j <= ub; ++j) {
                    res[id++] = matrix[j][i - j];
                }
            }
        }
        return res;
    }
};


#endif //LEETCODE_498_HPP
