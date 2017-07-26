
#include "447.hpp"
#include <gtest/gtest.h>

TEST(p447, example) {
    Solution s;
    vector<pair<int, int>> points = {{0, 0},
                                     {1, 0},
                                     {2, 0}};
    ASSERT_EQ(2, s.numberOfBoomerangs(points));
}
