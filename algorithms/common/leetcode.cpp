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
