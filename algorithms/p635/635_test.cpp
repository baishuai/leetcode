//
// Created by bai on 17-7-2.
//

#include "635.hpp"
#include <gtest/gtest.h>

TEST(p635, example) {
    LogSystem s;

    s.put(1, "2017:01:01:23:59:59");
    s.put(2, "2017:01:01:22:59:59");
    s.put(3, "2016:01:01:00:00:00");

    vector<int> exp = {1, 2, 3};
    vector<int> res = s.retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Year");
    sort(res.begin(), res.end());
    ASSERT_EQ(exp, res);

}