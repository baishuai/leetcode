//
// Created by bai on 17-6-30.
//

#include "238.hpp"
#include <gtest/gtest.h>

TEST(p238, example) {
    Solution s;
    vector<int> nums = {1, 2, 3, 4};

    vector<int> exp = {24, 12, 8, 6};
    ASSERT_EQ(exp, s.productExceptSelf(nums));
}