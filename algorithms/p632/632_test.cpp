//
// Created by bai on 17-7-2.
//


#include "632.hpp"
#include <gtest/gtest.h>

TEST(p632, example) {
    Solution s;

    vector<vector<int>> nums = {{4, 10, 15, 24, 26},
                                {0, 9,  12, 20},
                                {5, 18, 22, 30}};

    vector<int> res = {20, 24};
    ASSERT_EQ(res, s.smallestRange(nums));
}


TEST(p632, run1) {
    Solution s;

    vector<vector<int>> nums = {{-5, -4, -3, -2, -1, 1},
                                {1,  2,  3,  4,  5}};

    vector<int> res = {1, 1};
    ASSERT_EQ(res, s.smallestRange(nums));
}
