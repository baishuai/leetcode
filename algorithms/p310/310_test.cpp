//
// Created by bai on 17-6-30.
//

#include "310.hpp"
#include <gtest/gtest.h>

TEST(p310, example) {
    Solution s;

    vector<pair<int, int >> edges = {{1, 0},
                                     {1, 2},
                                     {1, 3}};

    vector<int> res = {1};
    ASSERT_EQ(res, s.findMinHeightTrees(4, edges));
}

TEST(p310, run1) {
    Solution s;

    vector<pair<int, int >> edges = {};

    vector<int> res = {0};
    ASSERT_EQ(res, s.findMinHeightTrees(1, edges));
}