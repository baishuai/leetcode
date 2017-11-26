
#include "729.hpp"
#include <gtest/gtest.h>

TEST(p729, example) {
    MyCalendar s;
    ASSERT_TRUE(s.book(10, 20));
    ASSERT_FALSE(s.book(15, 25));
    ASSERT_TRUE(s.book(20, 30));
}
