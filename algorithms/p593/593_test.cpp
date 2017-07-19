
#include "593.hpp"
#include <gtest/gtest.h>

TEST(p593, example) {
    Solution s;
    vector<int> p1 = {0, 0};
    vector<int> p2 = {1, 0};
    vector<int> p3 = {0, 1};
    vector<int> p4 = {1, 1};
    ASSERT_TRUE(s.validSquare(p1, p2, p3, p4));
}
