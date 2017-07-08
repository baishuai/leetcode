//
// Created by bai on 17-6-28.
//

#include "565.hpp"
#include <gtest/gtest.h>

TEST(p565, example) {
    Solution s;
    vector<int> arr = {5, 4, 0, 3, 1, 6, 2};
    ASSERT_EQ(4, s.arrayNesting(arr));
}