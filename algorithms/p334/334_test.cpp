
#include "334.hpp"
#include <gtest/gtest.h>

TEST(p334, example) {
    Solution s;

    vector<int> arr = {1, 2, 3, 4, 5};
    ASSERT_TRUE(s.increasingTriplet(arr));

    arr = {5, 4, 3, 2, 1};
    ASSERT_FALSE(s.increasingTriplet(arr));
}
