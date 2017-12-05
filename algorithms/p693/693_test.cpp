
#include "693.hpp"
#include <gtest/gtest.h>

TEST(p693, example) {
    Solution s;
    ASSERT_TRUE(s.hasAlternatingBits(5));
    ASSERT_TRUE(s.hasAlternatingBits(10));
    ASSERT_FALSE(s.hasAlternatingBits(7));
    ASSERT_FALSE(s.hasAlternatingBits(11));
}
