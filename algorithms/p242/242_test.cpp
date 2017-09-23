
#include "242.hpp"
#include <gtest/gtest.h>

TEST(p242, example) {
    Solution s;

    ASSERT_FALSE(s.isAnagram("rat", "cat"));
    ASSERT_TRUE(s.isAnagram("anagram", "nagaram"));
}
