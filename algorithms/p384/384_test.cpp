
#include "384.hpp"
#include <gtest/gtest.h>

TEST(p384, example) {
    vector<int> arr = {1, 2, 3, 4};
    Solution s(arr);
    ASSERT_EQ(arr, s.reset());
    auto a = s.shuffle();
    a = s.shuffle();
    ASSERT_EQ(arr, s.reset());
}
