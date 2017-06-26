//
// Created by bai on 17-6-26.
//

#include "313.hpp"
#include <gtest/gtest.h>

TEST(p313, example) {
    Solution s;
    vector<int> primes = {2, 3, 5};
//    ASSERT_EQ(1, s.nthSuperUglyNumber(1, primes));
//    ASSERT_EQ(2, s.nthSuperUglyNumber(2, primes));
    ASSERT_EQ(4, s.nthSuperUglyNumber(4, primes));
}