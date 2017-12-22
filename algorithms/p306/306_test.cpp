
#include "306.hpp"
#include <gtest/gtest.h>

TEST(p306, example) {
    Solution s;

    ASSERT_TRUE(s.isAdditiveNumber("112358"));
    ASSERT_TRUE(s.isAdditiveNumber("101"));
}
