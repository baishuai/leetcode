//
// Created by bai on 17-6-27.
//

#include "leetcode.hpp"

bool eqTree(TreeNode *lhs, TreeNode *rhs) {
    if (lhs == nullptr && rhs == nullptr) {
        return true;
    } else if (lhs == nullptr || rhs == nullptr) {
        return false;
    }
    return lhs->val == rhs->val && eqTree(lhs->left, rhs->left) && eqTree(lhs->right, rhs->right);
}

void releaseTree(TreeNode *root) {
    if (root == nullptr) {
        return;
    }
    releaseTree(root->left);
    releaseTree(root->right);
    delete root;
}


bool equalList(ListNode *l1, ListNode *l2) {
    while (l1 != nullptr && l2 != nullptr) {
        if (l1->val != l2->val) {
            return false;
        }
        l1 = l1->next;
        l2 = l2->next;
    }
    return (l1 == nullptr && l2 == nullptr);
}


void releaseList(ListNode *head, ListNode *entry) {
    if (entry != nullptr) {
        while (head != entry) {
            ListNode *p = head;
            head = head->next;
            p->next = nullptr;
            delete p;
        }
        head = entry->next;
        entry->next = nullptr;
    }
    while (head->next != nullptr) {
        ListNode *p = head;
        head = head->next;
        p->next = nullptr;
        delete p;
    }
    delete head;
}