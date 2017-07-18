
#include "503.hpp"
#include <gtest/gtest.h>

TEST(p503, example) {
    Solution s;
    vector<int> nums = {1, 2, 1};
    vector<int> exp = {2, -1, 2};
    ASSERT_EQ(exp, s.nextGreaterElements(nums));
}
