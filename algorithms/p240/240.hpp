
#ifndef LEETCODE_240_HPP
#define LEETCODE_240_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>

using namespace std;


/**
Write an efficient algorithm that searches for a value in an m x n matrix.
 This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
For example,

Consider the following matrix:

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
Given target = 5, return true.

Given target = 20, return false.


 */

class Solution {
public:
    bool searchMatrix(vector<vector<int>> &matrix, int target) {
        if (matrix.size() == 0 || matrix[0].size() == 0) {
            return false;
        }
        int i = 0, j = (int) (matrix[0].size() - 1);

        while (i < matrix.size() && j >= 0) {
            auto v = matrix[i][j];
            if (v == target) {
                break;
            } else if (v > target) {
                j--;
            } else {
                i++;
            }
        }

        return i < matrix.size() && j >= 0 && matrix[i][j] == target;
    }
};

#endif //LEETCODE_240_HPP
