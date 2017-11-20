
#include "443.hpp"
#include <gtest/gtest.h>

TEST(p443, example) {
    Solution s;
    vector<char> arr = {'a', 'a', 'b', 'b', 'c', 'c', 'c'};
    ASSERT_EQ(6, s.compress(arr));
}
