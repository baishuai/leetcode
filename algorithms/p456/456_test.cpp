//
// Created by bai on 17-6-25.
//

#include "456.hpp"
#include <gtest/gtest.h>


TEST(p456, example) {

    auto s = Solution();
    vector<int> res = {1, 2, 3, 4};
    ASSERT_EQ(false, s.find132pattern(res));
    vector<int> n1s = {3, 1, 4, 2};
    ASSERT_EQ(true, s.find132pattern(n1s));
    vector<int> n2s = {-1, 3, 2, 0};
    ASSERT_EQ(true, s.find132pattern(n2s));

    vector<int> n3s = {-2, 1, -2};
    ASSERT_EQ(false, s.find132pattern(n3s));

//    ASSERT_EQ(res, s.nextGreaterElement(n1s, n2s));
}