
#include "216.hpp"
#include <gtest/gtest.h>

TEST(p216, example) {
    Solution s;
    vector<vector<int>> a = {{1, 2, 4}};
    ASSERT_EQ(a, s.combinationSum3(3, 7));
}
