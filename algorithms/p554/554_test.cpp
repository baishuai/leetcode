
#include "554.hpp"
#include <gtest/gtest.h>

TEST(p554, example) {
    Solution s;
    
    vector<vector<int>> walls = {
            {1,2,2,1},
            {3,1,2},
            {1,3,2},
            {2,4},
            {3,1,2},
            {1,3,1,1},
    };
    ASSERT_EQ(2, s.leastBricks(walls));
}
