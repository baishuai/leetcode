//
// Created by bai on 17-6-30.
//

#ifndef LEETCODE_142_HPP
#define LEETCODE_142_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <numeric>

#include "../common/leetcode.hpp"

using namespace std;

/**
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

Note: Do not modify the linked list.

Follow up:
Can you solve it without using extra space?


 */

class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        if (head == nullptr || head->next == nullptr) {
            return nullptr;
        }
        ListNode *slow = head, *faster = head, *entry = head;

        while (faster->next != nullptr && faster->next->next != nullptr) {
            slow = slow->next;
            faster = faster->next->next;

            if (slow == faster) {
                while (entry != slow) {
                    slow = slow->next;
                    entry = entry->next;
                }
                return entry;
            }
        }
        return nullptr;
    }
};

#endif //LEETCODE_142_HPP
