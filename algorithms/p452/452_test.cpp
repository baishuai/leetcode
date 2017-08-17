
#include "452.hpp"
#include <gtest/gtest.h>

TEST(p452, example) {
    Solution s;
    vector<pair<int, int> > points = {{10, 16},
                                      {2,  8},
                                      {1,  6},
                                      {7,  12}};
    ASSERT_EQ(2, s.findMinArrowShots(points));
}
