
#include "328.hpp"
#include <gtest/gtest.h>


TEST(p328, example) {
    Solution s;

    ListNode *head = new ListNode(1);
    head->next = new ListNode(2);
    head->next->next = new ListNode(3);


    ListNode *node = new ListNode(1);
    node->next = new ListNode(3);
    node->next->next = new ListNode(2);

    head = s.oddEvenList(head);

    ASSERT_TRUE(equalList(head, node));
    releaseList(head);
    releaseList(node);
}
