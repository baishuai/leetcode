
#include "525.hpp"
#include <gtest/gtest.h>

TEST(p525, example) {
    Solution s;
    vector<int> arr = {0, 1};
    ASSERT_EQ(2, s.findMaxLength(arr));
}

TEST(p525, run1) {
    Solution s;
    vector<int> arr = {0, 0, 1, 0, 0, 0, 1, 1};
    ASSERT_EQ(6, s.findMaxLength(arr));
}
