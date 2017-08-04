
#ifndef LEETCODE_463_HPP
#define LEETCODE_463_HPP

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
    int islandPerimeter(vector<vector<int>> &grid) {
        int area = 0, connect = 0;
        for (int i = 0; i < grid.size(); ++i) {
            for (int j = 0; j < grid[i].size(); ++j) {
                area += grid[i][j];
                if (i < grid.size() - 1 && grid[i][j] == 1 && grid[i + 1][j] == 1)
                    connect++;
                if (j < grid[i].size() - 1 && grid[i][j] == 1 && grid[i][j + 1] == 1)
                    connect++;
            }
        }
        return area * 4 - connect * 2;
    }
};

#endif //LEETCODE_463_HPP
