
#include "643.hpp"
#include <gtest/gtest.h>

TEST(p643, example) {
    Solution s;

    vector<int> arr = {1, 12, -5, -6, 50, 3};
    ASSERT_EQ(12.75, s.findMaxAverage(arr, 4));
}
