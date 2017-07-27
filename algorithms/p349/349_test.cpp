
#include "349.hpp"
#include <gtest/gtest.h>

TEST(p349, example) {
    Solution s;
    vector<int> s1 = {1, 2, 2, 1};
    vector<int> s2 = {2, 2};
    vector<int> exp = {2};
    ASSERT_EQ(exp, s.intersection(s1, s2));
}
