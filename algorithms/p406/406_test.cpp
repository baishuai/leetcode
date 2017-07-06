//
// Created by bai on 17-6-30.
//

#include "406.hpp"
#include <gtest/gtest.h>

TEST(p406, example) {
    Solution s;

    vector<pair<int, int>> people = {{7, 0},
                                     {4, 4},
                                     {7, 1},
                                     {5, 0},
                                     {6, 1},
                                     {5, 2}};

    vector<pair<int, int>> exp = {{5, 0},
                                  {7, 0},
                                  {5, 2},
                                  {6, 1},
                                  {4, 4},
                                  {7, 1}};
    ASSERT_EQ(exp, s.reconstructQueue(people));

}
