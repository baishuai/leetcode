//
// Created by bai on 17-7-2.
//

#include "633.hpp"
#include <gtest/gtest.h>

TEST(p633, example) {
    Solution s;
    ASSERT_EQ(true, s.judgeSquareSum(5));
    ASSERT_EQ(false, s.judgeSquareSum(3));
}