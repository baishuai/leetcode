
#include "462.hpp"
#include <gtest/gtest.h>

TEST(p462, example) {
    Solution s;
    vector<int> arr = {4, 5, 1, 2};
    ASSERT_EQ(6, s.minMoves2(arr));
}
