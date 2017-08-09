
#include "599.hpp"
#include <gtest/gtest.h>

TEST(p599, example) {
    Solution s;
    vector<string> list1 = {"Shogun", "Tapioca Express", "Burger King", "KFC"};
    vector<string> list2 = {"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"};
    vector<string> res = {"Shogun"};
    ASSERT_EQ(res, s.findRestaurant(list1, list2));
}
