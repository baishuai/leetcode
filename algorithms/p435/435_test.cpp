
#include "435.hpp"
#include <gtest/gtest.h>

TEST(p435, example) {
    Solution s;
    vector<Interval> intervals = {Interval(1, 2), {2, 3}, {3, 4}, {1, 3}};
    ASSERT_EQ(1, s.eraseOverlapIntervals(intervals));
}
