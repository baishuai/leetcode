//
// Created by bai on 17-6-27.
//

#include "485.hpp"
#include <gtest/gtest.h>

TEST(p485, example) {
    Solution s;
    vector<int> arr = {1, 1, 0, 1, 1, 1};
    ASSERT_EQ(3, s.findMaxConsecutiveOnes(arr));
}