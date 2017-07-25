
#include "535.hpp"
#include <gtest/gtest.h>

TEST(p535, example) {
    Solution s;
    auto shortUrl = s.encode("leetcode.com");
    auto url = s.decode(shortUrl);
    ASSERT_EQ(url, "leetcode.com");
}
