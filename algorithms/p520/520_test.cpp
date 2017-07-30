
#include "520.hpp"
#include <gtest/gtest.h>

TEST(p520, example) {
    Solution s;
    ASSERT_EQ(true, s.detectCapitalUse("USA"));
    ASSERT_EQ(true, s.detectCapitalUse("Uss"));
    ASSERT_EQ(false, s.detectCapitalUse("UsA"));
}
