
#include "332.hpp"
#include <gtest/gtest.h>

TEST(p332, example) {
    Solution s;

    vector<pair<string, string>> tickets = {{"JFK", "KUL"},
                                            {"JFK", "NRT"},
                                            {"NRT", "JFK"}};

    vector<string> exp = {"JFK", "NRT", "JFK", "KUL"};

    ASSERT_EQ(exp, s.findItinerary(tickets));
}
