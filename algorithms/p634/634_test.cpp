//
// Created by bai on 17-7-2.
//

#include "634.hpp"
#include <gtest/gtest.h>

TEST(p634, example) {
    Solution s;

    ASSERT_EQ(1, s.findDerangement(2));
    ASSERT_EQ(2, s.findDerangement(3));
    ASSERT_EQ(9, s.findDerangement(4));
}

TEST(p634, run1) {
    Solution s;
    ASSERT_EQ(71100825, s.findDerangement(14));
}