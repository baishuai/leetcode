//
// Created by bai on 17-6-26.
//

#include "226.hpp"
#include <gtest/gtest.h>


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

TEST(p226, example) {
    Solution s;

    TreeNode *root = new TreeNode(1);
    root->left = new TreeNode(2);

    TreeNode *exp = new TreeNode(1);
    exp->right = new TreeNode(2);

    ASSERT_TRUE(eqTree(s.invertTree(root), exp));

    releaseTree(root);
    releaseTree(exp);

}
