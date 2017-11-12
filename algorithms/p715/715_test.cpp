
#include "715.hpp"
#include <gtest/gtest.h>

TEST(p715, example) {
    RangeModule r;
    r.addRange(10,20);
    r.removeRange(14,16);
    ASSERT_TRUE(r.queryRange(10,14));
    ASSERT_FALSE(r.queryRange(13,15));
}
