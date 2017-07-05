
#ifndef LEETCODE_234_HPP
#define LEETCODE_234_HPP

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
 * Given a singly linked list, determine if it is a palindrome.

Follow up:
Could you do it in O(n) time and O(1) space?
 */




class Solution {


public:
    bool isPalindrome(ListNode *head) {
        if (head == nullptr || head->next == nullptr) {
            return true;
        }
        ListNode *slow = head;
        ListNode *fast = head;
        while (fast->next != nullptr && fast->next->next != nullptr) {
            slow = slow->next;
            fast = fast->next->next;
        }
        slow->next = reverseList(slow->next);
        slow = slow->next;
        while (slow != nullptr) {
            if (head->val != slow->val) {
                return false;
            }
            head = head->next;
            slow = slow->next;
        }
        return true;
    }

private:
    ListNode *reverseList(ListNode *head) {
        ListNode *guard = nullptr, *tmp = nullptr;
        while (head != nullptr) {
            tmp = head;
            head = head->next;
            tmp->next = guard;
            guard = tmp;
        }
        return guard;
    }
};


#endif //LEETCODE_234_HPP
