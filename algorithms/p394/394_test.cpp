//
// Created by bai on 17-6-30.
//

#include "394.hpp"
#include <gtest/gtest.h>

TEST(p394, example) {
    Solution s;

    ASSERT_EQ("aaabcbc", s.decodeString("3[a]2[bc]"));
    ASSERT_EQ("accaccacc", s.decodeString("3[a2[c]]"));
    ASSERT_EQ("abcabccdcdcdef", s.decodeString("2[abc]3[cd]ef"));
}

TEST(p394, run2) {
    Solution s;
    ASSERT_EQ("aabbbbbbaabbbbbb", s.decodeString("2[aa3[bb]]1[]"));
}