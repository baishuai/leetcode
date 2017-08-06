
#include "240.hpp"
#include <gtest/gtest.h>

TEST(p240, example) {
    Solution s;

    vector<vector<int>> mat = {
            {1,  4,  7,  11, 15},
            {2,  5,  8,  12, 19},
            {3,  6,  9,  16, 22},
            {10, 13, 14, 17, 24},
            {18, 21, 23, 26, 30}
    };

    ASSERT_TRUE(s.searchMatrix(mat, 5));
    ASSERT_TRUE(s.searchMatrix(mat, 18));
    ASSERT_FALSE(s.searchMatrix(mat, 20));
    ASSERT_FALSE(s.searchMatrix(mat, 25));
}

TEST(p240, run1) {
    Solution s;
    vector<vector<int>> mat = {{-1},
                               {-1}};
    ASSERT_FALSE(s.searchMatrix(mat, -2));
}
