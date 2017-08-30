
#include "455.hpp"
#include <gtest/gtest.h>

TEST(p455, example) {
    Solution s;
    vector<int> g = {1, 2, 3};
    vector<int> c = {1, 1};
    ASSERT_EQ(1, s.findContentChildren(g, c));
}

TEST(p455, run1) {
    Solution s;
    vector<int> g = {1, 2, 3};
    vector<int> c = {3};
    ASSERT_EQ(1, s.findContentChildren(g, c));
}