//
// Created by bai on 17-6-25.
//

#include "496.hpp"
#include <gtest/gtest.h>

TEST(p496, example) {

    auto s = Solution();
    vector<int> res = {-1, 3, -1};
    vector<int> n1s = {4, 1, 2};
    vector<int> n2s = {1, 3, 4, 2};

    ASSERT_EQ(res, s.nextGreaterElement(n1s, n2s));
}

TEST(p496, run1) {

    auto s = Solution();
    vector<int> res = {3, -1};
    vector<int> n1s = {2, 4};
    vector<int> n2s = {1, 2, 3, 4};

    ASSERT_EQ(res, s.nextGreaterElement(n1s, n2s));
}