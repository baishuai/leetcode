
#include "436.hpp"
#include <gtest/gtest.h>

TEST(p436, example) {
    Solution s;
    vector<Interval> intervals = {Interval{1, 2}};
    vector<int> exp = {-1};
    ASSERT_EQ(exp, s.findRightInterval(intervals));
}


TEST(p436, example1) {
    Solution s;
    vector<Interval> intervals = {Interval{3, 4}, Interval{2, 3}, Interval{1, 2}};
    vector<int> exp = {-1, 0, 1};
    ASSERT_EQ(exp, s.findRightInterval(intervals));
}
