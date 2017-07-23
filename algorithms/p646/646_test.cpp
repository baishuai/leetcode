
#include "646.hpp"
#include <gtest/gtest.h>

TEST(p646, example) {
    Solution s;
    vector<vector<int>> pairs = {{1, 6},
                                 {2, 3},
                                 {3, 4},
                                 {4, 5},
                                 {7, 8}};
    ASSERT_EQ(3, s.findLongestChain(pairs));
}
