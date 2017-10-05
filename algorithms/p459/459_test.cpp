
#include "459.hpp"
#include <gtest/gtest.h>

TEST(p459, example) {
    Solution s;
    ASSERT_EQ(false, s.repeatedSubstringPattern("aba"));
    ASSERT_EQ(true, s.repeatedSubstringPattern("abab"));
}
