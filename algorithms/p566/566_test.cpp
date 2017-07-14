
#include "566.hpp"
#include <gtest/gtest.h>

TEST(p566, example) {
    Solution s;
    vector<vector<int>> nums = {{1, 2},
                                {3, 4}};
    vector<vector<int> > exp = {{1, 2, 3, 4}};

    ASSERT_EQ(exp, s.matrixReshape(nums, 1, 4));
    ASSERT_EQ(nums, s.matrixReshape(nums, 2, 4));
}
