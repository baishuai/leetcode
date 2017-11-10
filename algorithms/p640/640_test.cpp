
#include "640.hpp"
#include <gtest/gtest.h>

TEST(p640, example) {
    Solution s;
    ASSERT_EQ("x=2", s.solveEquation("x+5-3+x=6+x-2"));
}


TEST(p640, example1) {
    Solution s;
    ASSERT_EQ("x=-1", s.solveEquation("2x+3x-6x=x+2"));
}



TEST(p640, example2) {
    Solution s;
    ASSERT_EQ("x=-2", s.solveEquation("2=-x"));
}
