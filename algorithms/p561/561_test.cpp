//
// Created by bai on 17-6-26.
//

#include "561.hpp"
#include <gtest/gtest.h>

TEST(p561, example) {
    Solution s;
    vector<int> nums = {1, 4, 3, 2};
    ASSERT_EQ(4, s.arrayPairSum(nums));
}