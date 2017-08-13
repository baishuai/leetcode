
#ifndef LEETCODE_328_HPP
#define LEETCODE_328_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>
#include "../common/leetcode.hpp"

using namespace std;

/**
Given a singly linked list, group all odd nodes together followed by the even nodes.
 Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

Example:
Given 1->2->3->4->5->NULL,
return 1->3->5->2->4->NULL.

Note:
The relative order inside both the even and odd groups should remain as it was in the input.
The first node is considered odd, the second node even and so on ...
 */

class Solution {
public:
    ListNode *oddEvenList(ListNode *head) {
        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        ListNode *oddhead = head, *evenHead = head->next;

        ListNode *odd = oddhead, *even = evenHead, *cur = evenHead->next;
        while (cur != nullptr) {
            odd->next = cur;
            odd = odd->next;
            cur = cur->next;
            if (cur != nullptr) {
                even->next = cur;
                even = even->next;
                cur = cur->next;
            }
        }
        odd->next = evenHead;
        even->next = nullptr;
        return oddhead;
    }
};


#endif //LEETCODE_328_HPP
