
#include "318.hpp"
#include <gtest/gtest.h>

TEST(p318, example) {
    Solution s;

    vector<string> arr = {"abcw", "baz", "foo", "bar", "xtfn", "abcdef"};
    ASSERT_EQ(16, s.maxProduct(arr));
}
