//
// Created by bai on 17-6-28.
//


#include "560.hpp"
#include <gtest/gtest.h>

TEST(p560, example) {
    Solution s;
    vector<int> arr = {1, 1, 1};
    ASSERT_EQ(2, s.subarraySum(arr, 2));
}

TEST(p560, run1) {
    Solution s;
    vector<int> arr = {0, 0, 0, 0, 0, 0, 0, 0, 0, 0};
    ASSERT_EQ(55, s.subarraySum(arr, 0));
}