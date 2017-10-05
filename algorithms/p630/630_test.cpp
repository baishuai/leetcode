//
// Created by bai on 17-6-25.
//

#include "630.hpp"
#include <gtest/gtest.h>


TEST(p630, example) {
    auto s = Solution();
    vector<vector<int>> course = {{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}};
    ASSERT_EQ(3, s.scheduleCourse(course));
}



