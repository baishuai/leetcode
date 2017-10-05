
#include "670.hpp"
#include <gtest/gtest.h>

TEST(p670, example) {
    Solution s;
    ASSERT_EQ(7236, s.maximumSwap(2736));
    ASSERT_EQ(9913, s.maximumSwap(1993));
}


TEST(p670, run1){
    Solution s;
    ASSERT_EQ(98863, s.maximumSwap(98368));
}