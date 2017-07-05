
#include "234.hpp"
#include <gtest/gtest.h>

TEST(p234, example) {
    Solution s;

    ListNode *h = new ListNode(1);
    ASSERT_EQ(true, s.isPalindrome(h));
    h->next = new ListNode(2);
    ASSERT_EQ(false, s.isPalindrome(h));
    h->next->val = 1;
    ASSERT_EQ(true, s.isPalindrome(h));
    h->next->val = 2;
    h->next->next = new ListNode(1);
    ASSERT_EQ(true, s.isPalindrome(h));

    releaseList(h);
}

