
#include "424.hpp"
#include <gtest/gtest.h>

TEST(p424, example) {
    Solution s;
    ASSERT_EQ(4, s.characterReplacement("AABABBA",1));
    ASSERT_EQ(4, s.characterReplacement("AABA",1));
}
