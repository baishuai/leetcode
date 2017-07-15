//
// Created by bai on 17-6-27.
//

#ifndef LEETCODE_LEETCODE_HPP
#define LEETCODE_LEETCODE_HPP

#include <iostream>
#include <vector>

using namespace std;

//Definition for a binary tree node.
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;

    explicit TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};


bool eqTree(TreeNode *lhs, TreeNode *rhs);

void releaseTree(TreeNode *root);

struct ListNode {
    int val;
    ListNode *next;

    explicit ListNode(int x) : val(x), next(NULL) {}
};


bool equalList(ListNode *l1, ListNode *l2);

void releaseList(ListNode *head, ListNode *entry = nullptr);

#endif //LEETCODE_LEETCODE_HPP
