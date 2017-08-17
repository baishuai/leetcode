
#ifndef LEETCODE_382_HPP
#define LEETCODE_382_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>
#include <stack>
#include <string>
#include "../common/leetcode.hpp"

using namespace std;

/*
Given a singly linked list, return a random node's value from the linked list.
 Each node must have the same probability of being chosen.

Follow up:
What if the linked list is extremely large and its length is unknown to you?
 Could you solve this efficiently without using extra space?

Example:

// Init a singly linked list [1,2,3].
ListNode head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
Solution solution = new Solution(head);

// getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
solution.getRandom();
 */


class Solution {
public:
    /** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
    Solution(ListNode *head) : head(head) {}

    /** Returns a random node's value. */
    int getRandom() {
        int res = head->val;
        ListNode *node = head->next;
        int i = 2;
        while (node) {
            int j = rand() % i;
            if (j == 0)
                res = node->val;
            i++;
            node = node->next;
        }
        return res;
    }

private:
    ListNode *head;
};

#endif //LEETCODE_382_HPP
