
#include "345.hpp"
#include <gtest/gtest.h>

TEST(p345, example) {
    Solution s;
    ASSERT_EQ("holle", s.reverseVowels("hello"));
    ASSERT_EQ("leotcede", s.reverseVowels("leetcode"));
}
