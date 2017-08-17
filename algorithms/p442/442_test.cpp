
#include "442.hpp"
#include <gtest/gtest.h>

TEST(p442, example) {
    Solution s;
    vector<int> arr = {4, 3, 2, 7, 8, 2, 3, 1};
    vector<int> exp = {2, 3};
    ASSERT_EQ(exp, s.findDuplicates(arr));
}
