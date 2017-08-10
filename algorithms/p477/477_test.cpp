
#include "477.hpp"
#include <gtest/gtest.h>

TEST(p477, example) {
    Solution s;
    vector<int> nums = {4, 14, 2};
    ASSERT_EQ(6, s.totalHammingDistance(nums));
}
