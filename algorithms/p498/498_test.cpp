//
// Created by bai on 17-6-27.
//

#include "498.hpp"
#include <gtest/gtest.h>

TEST(p498, example) {
    Solution s;

    vector<vector<int>> mat = {{1, 2, 3},
                               {4, 5, 6},
                               {7, 8, 9}};

    vector<int> exp = {1, 2, 4, 7, 5, 3, 6, 8, 9};
    ASSERT_EQ(exp, s.findDiagonalOrder(mat));
}

TEST(p498, run1) {
    Solution s;

    vector<vector<int>> mat = {{2, 3}};

    vector<int> exp = {2, 3};
    ASSERT_EQ(exp, s.findDiagonalOrder(mat));

    mat = {{2},
           {3}};

    ASSERT_EQ(exp, s.findDiagonalOrder(mat));
}