
#include "540.hpp"
#include <gtest/gtest.h>

TEST(p540, example) {
    Solution s;
    vector<int> nums = {0, 1, 1};
    ASSERT_EQ(0, s.singleNonDuplicate(nums));
}
