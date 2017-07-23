
#include "645.hpp"
#include <gtest/gtest.h>

TEST(p645, example) {
    Solution s;
    vector<int> nums = {2, 2};
    vector<int> exp = {2, 1};
    auto ans = s.findErrorNums(nums);
    ASSERT_EQ(ans, exp);
}
