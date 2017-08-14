
#include "386.hpp"
#include <gtest/gtest.h>

TEST(p386, example) {
    Solution s;
    vector<int> as = {1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9};
    ASSERT_EQ(as, s.lexicalOrder(13));
}