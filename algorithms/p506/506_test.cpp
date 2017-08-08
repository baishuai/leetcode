
#include "506.hpp"
#include <gtest/gtest.h>

TEST(p506, example) {
    Solution s;
    vector<string> result = {"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"};
    vector<int> nums = {5, 4, 3, 2, 1};
    ASSERT_EQ(result, s.findRelativeRanks(nums));
}
