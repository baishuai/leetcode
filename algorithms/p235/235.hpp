//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_235_HPP
#define LEETCODE_235_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

#include "../common/leetcode.hpp"

using namespace std;

class Solution {
public:

    TreeNode *lowestCommonAncestor(TreeNode *root, TreeNode *p, TreeNode *q) {
        if (root == nullptr || p == nullptr || q == nullptr) {
            return nullptr;
        }
        while ((root->val - p->val) * (root->val - q->val) > 0) {
            root = p->val < root->val ? root->left : root->right;
        }
        return root;
    }
};

#endif //LEETCODE_235_HPP
