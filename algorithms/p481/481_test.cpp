
#include "481.hpp"
#include <gtest/gtest.h>

TEST(p481, example) {
    Solution s;
    ASSERT_EQ(2, s.magicalString(4));
}

TEST(p481, run1) {
    Solution s;
    ASSERT_EQ(4, s.magicalString(7));
}
