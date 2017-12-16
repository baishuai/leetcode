
#include "735.hpp"
#include <gtest/gtest.h>

TEST(p735, example) {
    Solution s;

    vector<int> asteroids = {5, 10, -5};
    vector<int> exp = {5, 10};
    ASSERT_EQ(exp, s.asteroidCollision(asteroids));
}
