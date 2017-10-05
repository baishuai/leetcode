
#include "518.hpp"
#include <gtest/gtest.h>

TEST(p518, example) {
    Solution s;
    vector<int> coins = {1,2,5};
    ASSERT_EQ(4, s.change(5, coins));
}
