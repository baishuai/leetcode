
#include "495.hpp"
#include <gtest/gtest.h>

TEST(p495, example) {
    Solution s;
    vector<int> arr = {1, 2};
    ASSERT_EQ(s.findPoisonedDuration(arr, 2), 3);
}
