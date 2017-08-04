
#include "463.hpp"
#include <gtest/gtest.h>

TEST(p463, example) {
    Solution s;
    vector<vector<int>> grid = {
            {0, 1, 0, 0},
            {1, 1, 1, 0},
            {0, 1, 0, 0},
            {1, 1, 0, 0}
    };
    ASSERT_EQ(16, s.islandPerimeter(grid));
}
