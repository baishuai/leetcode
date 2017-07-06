//
// Created by bai on 17-6-28.
//


#include "583.hpp"
#include <gtest/gtest.h>

TEST(p583, example) {
    Solution s;

    ASSERT_EQ(2, s.minDistance("sea", "eat"));
    ASSERT_EQ(2, s.minDistance("se", "ea"));
}
