
#include "665.hpp"
#include <gtest/gtest.h>

TEST(p665, example) {
    Solution s;
    vector<int> arr1 = {1, 2, 3, 4};
    ASSERT_TRUE(s.checkPossibility(arr1));
    vector<int> arr2 = {3, 2, 1};
    ASSERT_FALSE(s.checkPossibility(arr2));
}
