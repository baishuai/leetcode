//
// Created by bai on 17-6-28.
//

#include "287.hpp"
#include <gtest/gtest.h>

TEST(p287, example) {
    Solution s;

    vector<int> arr = {1, 4, 2, 2, 3};
    ASSERT_EQ(2, s.findDuplicate(arr));

    arr = {2, 1, 1};
    ASSERT_EQ(1, s.findDuplicate(arr));
}


TEST(p287, run1) {
    Solution s;
    vector<int> arr = {2, 1, 1};
    ASSERT_EQ(1, s.findDuplicate(arr));
}


TEST(p287, run2) {
    Solution s;
    vector<int> arr = {1, 1, 2};
    ASSERT_EQ(1, s.findDuplicate(arr));
}

TEST(p287, run3) {
    Solution s;
    vector<int> arr = {1, 3, 4, 2, 2};
    ASSERT_EQ(2, s.findDuplicate(arr));
}