
#include "260.hpp"
#include <gtest/gtest.h>

TEST(p260, example) {
    Solution s;
    vector<int> arr = {1, 2, 3, 1, 2, 5};
    vector<int> exp = {3, 5};
    auto r = s.singleNumber(arr);
    sort(r.begin(), r.end());
    ASSERT_EQ(exp, r);
}
