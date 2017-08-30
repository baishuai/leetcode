
#ifndef LEETCODE_598_HPP
#define LEETCODE_598_HPP

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

/*
Given an m * n matrix M initialized with all 0's and several update operations.

Operations are represented by a 2D array, and each operation is represented by
 an array with two positive integers a and b,
 which means M[i][j] should be added by one for all 0 <= i < a and 0 <= j < b.

You need to count and return the number of maximum integers in the matrix after performing all the operations.
 */


class Solution {
public:
    int maxCount(int m, int n, vector<vector<int>> &ops) {
        int row = m;
        int col = n;
        for (const auto &p : ops) {
            row = min(row, p[0]);
            col = min(col, p[1]);
        }
        return row * col;
    }
};

#endif //LEETCODE_598_HPP
