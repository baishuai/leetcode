#include "631.hpp"
#include <gtest/gtest.h>


TEST(p631, example) {
    Excel e = Excel(3, 'C');

    e.set(1, 'A', 2);

    ASSERT_EQ(4, e.sum(3, 'C', {"A1", "A1:B2"}));
    e.set(2, 'B', 2);
}

TEST(p631, run1) {
    Excel e = Excel(3, 'C');
    ASSERT_EQ(0, e.sum(1, 'A', {"A2"}));
    e.set(2, 'A', 1);
    ASSERT_EQ(1, e.sum(3, 'A', {"A1"}));
}

TEST(p631, run2) {
    Excel e = Excel(26, 'Z');
    e.set(1, 'A', 1);
    e.set(1, 'I', 1);

    ASSERT_EQ(3, e.sum(7, 'D', {"A1:D6", "A1:G3", "A1:C12"}));
    ASSERT_EQ(11, e.sum(10, 'G', {"A1:D7", "D1:F10", "D3:I8", "I1:I9"}));
}

TEST(p631, run3) {
    Excel e = Excel(3, 'C');
    ASSERT_EQ(0, e.sum(1, 'A', {"A2"}));
    e.set(2, 'A', 1);
    ASSERT_EQ(1, e.get(1, 'A'));
}