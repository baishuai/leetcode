//
// Created by bai on 17-6-28.
//

#include "367.hpp"
#include <gtest/gtest.h>

TEST(p367, example) {
    Solution s;
    for (int i = 1; i < 100; ++i) {
        ASSERT_TRUE(s.isPerfectSquare(i * i));
        ASSERT_FALSE(s.isPerfectSquare(i * i + 1));
    }
}