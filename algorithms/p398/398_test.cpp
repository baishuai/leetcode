
#include "398.hpp"
#include <gtest/gtest.h>

TEST(p398, example) {
    vector<int> nums = {1};
    Solution s(nums);
    ASSERT_EQ(0, s.pick(1));
}
