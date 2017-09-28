
#include "457.hpp"
#include <gtest/gtest.h>

TEST(p457, example) {
    Solution s;
    vector<int> arr = {2, -1, 1, 2, 2};
    ASSERT_TRUE(s.circularArrayLoop(arr));
}


TEST(p457, example_false) {
    Solution s;
    vector<int> arr = {-1, 2};
    ASSERT_FALSE(s.circularArrayLoop(arr));
}
