
#ifndef LEETCODE_237_HPP
#define LEETCODE_237_HPP

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

class Solution {
public:
    void deleteNode(ListNode *node) {
        if (node == nullptr || node->next == nullptr)
            return;
        ListNode *nn = node->next->next;
        while (nn != nullptr) {
            node->val = node->next->val;
            node = node->next;
            nn = nn->next;
        }
        node->val = node->next->val;
        delete node->next;
        node->next = nullptr;
    }
};

#endif //LEETCODE_237_HPP
