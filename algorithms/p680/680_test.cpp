
#include "680.hpp"
#include <gtest/gtest.h>

TEST(p680, example) {
    Solution s;
    ASSERT_TRUE(s.validPalindrome("aba"));
    ASSERT_TRUE(s.validPalindrome("abca"));
}
