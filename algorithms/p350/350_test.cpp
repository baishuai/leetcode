
#include "350.hpp"
#include <gtest/gtest.h>

TEST(p350, example) {
    Solution s;
    vector<int> s1 = {1, 2, 2, 1};
    vector<int> s2 = {2, 2};
    vector<int> exp = {2, 2};
    ASSERT_EQ(exp, s.intersect(s1, s2));
}
