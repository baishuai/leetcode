
#include "556.hpp"
#include <gtest/gtest.h>

TEST(p556, example) {
    Solution s;
    ASSERT_EQ(-1, s.nextGreaterElement(21));
    ASSERT_EQ(21, s.nextGreaterElement(12));
}
