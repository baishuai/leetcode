
#include "445.hpp"
#include <gtest/gtest.h>

TEST(p445, example) {
    Solution s;
    ListNode *l1 = new ListNode(5);
    ListNode *l2 = new ListNode(5);
    ListNode *exp = new ListNode(1);
    exp->next = new ListNode(0);
    ListNode *ans = s.addTwoNumbers(l1, l2);
    ASSERT_TRUE(equalList(exp, ans));
    releaseList(l1);
    releaseList(l2);
    releaseList(exp);
    releaseList(ans);
}
