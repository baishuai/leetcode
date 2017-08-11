
#include "524.hpp"
#include <gtest/gtest.h>

TEST(p524, example) {
    Solution s;
    vector<string> d = {"ale", "apple", "monkey", "plea"};
    string str = "abpcplea";
    ASSERT_EQ("apple", s.findLongestWord(str, d));
}
