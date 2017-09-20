
#ifndef LEETCODE_661_HPP
#define LEETCODE_661_HPP

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
    vector<vector<int>> imageSmoother(vector<vector<int>> &M) {
        auto m = M.size();
        if (m == 0)
            return {};
        auto n = M[0].size();
        vector<vector<int>> dirs = {{0,  1},
                                    {0,  -1},
                                    {1,  0},
                                    {-1, 0},
                                    {-1, -1},
                                    {1,  1},
                                    {-1, 1},
                                    {1,  -1}};
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                int sum = M[i][j], cnt = 1;
                for (int k = 0; k < dirs.size(); k++) {
                    int x = i + dirs[k][0], y = j + dirs[k][1];
                    if (x < 0 || x > m - 1 || y < 0 || y > n - 1)
                        continue;
                    sum += (M[x][y] & 0xFF);
                    cnt++;
                }

                M[i][j] |= ((sum / cnt) << 8);

            }
        }

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                M[i][j] >>= 8;
            }
        }
        return M;
    }
};


#endif //LEETCODE_661_HPP
