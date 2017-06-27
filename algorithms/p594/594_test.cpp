//
// Created by bai on 17-6-26.
//



#include "594.hpp"
#include <gtest/gtest.h>


TEST(p594, example) {
    auto s = Solution();

    vector<int> in = {1, 3, 2, 2, 5, 2, 3, 7};
    ASSERT_EQ(5, s.findLHS(in));
}


TEST(p594, t1) {
    auto s = Solution();
    vector<int> in = {1, 3};
    ASSERT_EQ(0, s.findLHS(in));
}


TEST(p594, t2) {
    auto s = Solution();
    vector<int> in = {1, 2, 3, 3, 1, -14, 13, 4};
    ASSERT_EQ(3, s.findLHS(in));
}