
#include "412.hpp"
#include <gtest/gtest.h>

TEST(p412, example) {
    Solution s;
    vector<string> exp = {
            "1",
            "2",
            "Fizz",
            "4",
            "Buzz",
            "Fizz",
            "7",
            "8",
            "Fizz",
            "Buzz",
            "11",
            "Fizz",
            "13",
            "14",
            "FizzBuzz"
    };
    ASSERT_EQ(exp, s.fizzBuzz(15));
}
