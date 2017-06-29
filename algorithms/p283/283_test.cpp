//
// Created by bai on 17-6-28.
//

#include "283.hpp"
#include <gtest/gtest.h>

TEST(p283, example) {
    Solution s;

    vector<int> arr = {0, 1, 0, 3, 12};
    vector<int> exp = {1, 3, 12, 0, 0};
    s.moveZeroes(arr);
    ASSERT_EQ(exp, arr);
}