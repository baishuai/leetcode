//
// Created by bai on 17-6-30.
//

#include "142.hpp"
#include <gtest/gtest.h>

TEST(p142, example) {
    Solution s;


    ListNode *head = new ListNode(1);
    head->next = new ListNode(2);
    head->next->next = new ListNode(3);
    head->next->next->next = new ListNode(4);
    head->next->next->next->next = head->next;

    ASSERT_EQ(head->next, s.detectCycle(head));
    ASSERT_EQ(2, s.detectCycle(head)->val);

    releaseList(head, head->next);
}