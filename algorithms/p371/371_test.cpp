//
// Created by bai on 17-6-28.
//

#include "371.hpp"
#include <gtest/gtest.h>

TEST(p371, example) {
    Solution s;

    for (int i = 0; i < 1000; ++i) {
        int a = rand();
        int b = rand();
        ASSERT_EQ(a+b, s.getSum(a,b));
    }
}