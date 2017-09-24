
#include "667.hpp"
#include <gtest/gtest.h>

TEST(p667, example) {
    Solution s;
    vector<int> arr = {1, 3, 2};
    ASSERT_EQ(arr, s.constructArray(3, 2));
}
