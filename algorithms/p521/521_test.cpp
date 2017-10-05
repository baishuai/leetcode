
#include "521.hpp"
#include <gtest/gtest.h>

TEST(p521, example) {
    Solution s;
    ASSERT_EQ(3, s.findLUSlength("abc", "ebc"));
}
