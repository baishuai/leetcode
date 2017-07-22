
#ifndef LEETCODE_445_HPP
#define LEETCODE_445_HPP

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
You are given two non-empty linked lists representing two non-negative integers.
 The most significant digit comes first and each of their nodes contain a single digit.
 Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:
What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

Example:

Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 8 -> 0 -> 7
 */

class Solution {
public:
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2) {
        stack<int> s1, s2;
        while (l1 != nullptr) {
            s1.emplace(l1->val);
            l1 = l1->next;
        }
        while (l2 != nullptr) {
            s2.emplace(l2->val);
            l2 = l2->next;
        }
        int carry = 0;
        ListNode *head = nullptr;
        while (!s1.empty() && !s2.empty()) {
            ListNode *t = new ListNode(s1.top() + s2.top() + carry);
            carry = t->val / 10;
            t->val = t->val % 10;
            t->next = head;
            head = t;
            s1.pop();
            s2.pop();
        }
        stack<int> &s = s1.empty() ? s2 : s1;
        while (!s.empty()) {
            ListNode *t = new ListNode(s.top() + carry);
            carry = t->val / 10;
            t->val = t->val % 10;
            t->next = head;
            head = t;
            s.pop();
        }
        if (carry != 0) {
            ListNode *t = new ListNode(carry);
            t->next = head;
            head = t;
        }
        return head;
    }

};


#endif //LEETCODE_445_HPP
