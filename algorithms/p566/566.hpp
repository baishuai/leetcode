
#ifndef LEETCODE_566_HPP
#define LEETCODE_566_HPP

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
You're given a matrix represented by a two-dimensional array,
 and two positive integers r and c representing the row number
 and column number of the wanted reshaped matrix, respectively.

The reshaped matrix need to be filled with all the elements of the original matrix
 in the same row-traversing order as they were.

If the 'reshape' operation with given parameters is possible and legal,
 output the new reshaped matrix; Otherwise, output the original matrix.
 */


class Solution {
public:
    vector<vector<int>> matrixReshape(vector<vector<int>> &nums, int r, int c) {
        auto m = nums.size();
        if (m == 0) {
            return nums;
        }
        auto n = nums[0].size();

        if (m == r || m * n != r * c) {
            return nums;
        }

        vector<vector<int>> res((unsigned long) r, vector<int>((unsigned long) c));
        int pi = 0, pj = 0;
        for (auto row : nums) {
            for (auto v : row) {
                res[pi][pj++] = v;
                if (pj == c) {
                    pj = 0;
                    pi++;
                }
                if (pi == r) {
                    return res;
                }
            }
        }
        return res;
    }
};


#endif //LEETCODE_566_HPP
