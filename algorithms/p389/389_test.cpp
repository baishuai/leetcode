//
// Created by bai on 17-6-28.
//

#include "389.hpp"
#include <gtest/gtest.h>

TEST(p389, example) {
    Solution s;

    ASSERT_EQ('e', s.findTheDifference("abcd", "abcde"));
}