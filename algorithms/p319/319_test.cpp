
#include "319.hpp"
#include <gtest/gtest.h>

TEST(p319, example) {
    Solution s;
    ASSERT_EQ(1, s.bulbSwitch(3));
    ASSERT_EQ(2, s.bulbSwitch(4));
}
