#include "629.hpp"
#include <gtest/gtest.h>

TEST(p629, example){
	auto s = Solution();

	ASSERT_EQ(1, s.kInversePairs(3,0));
	ASSERT_EQ(2, s.kInversePairs(3,1));

    ASSERT_EQ(1, s.kInversePairs2(3,0));
    ASSERT_EQ(2, s.kInversePairs2(3,1));
}