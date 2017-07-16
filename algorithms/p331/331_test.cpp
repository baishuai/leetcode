
#include "331.hpp"
#include <gtest/gtest.h>

TEST(p331, example) {
    Solution s;

    string tree = "9,3,4,#,#,1,#,#,2,#,6,#,#";

    ASSERT_TRUE(s.isValidSerialization(tree));
}
