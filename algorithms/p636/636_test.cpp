
#include "636.hpp"
#include <gtest/gtest.h>

TEST(p636, example) {
    Solution s;

    vector<int> exp = {3, 4};
    vector<string> logs = {"0:start:0", "1:start:2", "1:end:5", "0:end:6"};
    ASSERT_EQ(exp, s.exclusiveTime(2, logs));
}
