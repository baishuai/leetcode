
#include "378.hpp"
#include <gtest/gtest.h>

TEST(p378, example) {
    Solution s;
    vector<vector<int>> mat = {{1,  5,  9},
                               {10, 11, 12},
                               {12, 13, 15}};
    ASSERT_EQ(13, s.kthSmallest(mat, 8));

}
