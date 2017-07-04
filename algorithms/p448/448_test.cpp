//
// Created by bai on 17-6-27.
//

#include "448.hpp"
#include <gtest/gtest.h>

TEST(p448, example) {
    Solution s;
    vector<int> arr = {4, 3, 2, 7, 8, 2, 3, 1};

    vector<int> ans = {5, 6};
    ASSERT_EQ(ans, s.findDisappearedNumbers(arr));
}