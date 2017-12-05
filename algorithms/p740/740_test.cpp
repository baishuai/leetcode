
#include "740.hpp"
#include <gtest/gtest.h>

TEST(p740, example) {
    Solution s;
    vector<int > arr = {2,2,3,3,3,4};
    ASSERT_EQ(9, s.deleteAndEarn(arr));
}
