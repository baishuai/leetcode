//
// Created by bai on 17-6-27.
//

#include "450.hpp"
#include <gtest/gtest.h>

TEST(p450, example) {
    Solution s;

    TreeNode *root = new TreeNode(5);
    root->left = new TreeNode(3);
    root->right = new TreeNode(6);
    root->left->left = new TreeNode(2);
    root->left->right = new TreeNode(4);
    root->right->right = new TreeNode(7);

    TreeNode *exp = new TreeNode(5);
    exp->left = new TreeNode(4);
    exp->right = new TreeNode(6);
    exp->left->left = new TreeNode(2);
    exp->right->right = new TreeNode(7);

    ASSERT_TRUE(eqTree(exp, s.deleteNode(root, 3)));

    releaseTree(root);
    releaseTree(exp);
}