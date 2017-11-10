
#include "713.hpp"
#include <gtest/gtest.h>

TEST(p713, example) {
    Solution s;
    vector<int> arr = {10, 5, 2, 6};
    ASSERT_EQ(8, s.numSubarrayProductLessThanK(arr, 100));
}


TEST(p713, example2) {
    Solution s;
    vector<int> arr = {10, 5, 2, 6};
    ASSERT_EQ(0, s.numSubarrayProductLessThanK(arr, 0));
}
