//
// Created by bai on 17-6-28.
//

#include "513.hpp"
#include <gtest/gtest.h>

TEST(p513, example) {
    Solution s;
    TreeNode *root = new TreeNode(1);
    root->right = new TreeNode(3);
    root->left = new TreeNode(2);
    root->left->right = new TreeNode(4);
    root->right->right = new TreeNode(6);
    ASSERT_EQ(4, s.findBottomLeftValue(root));
    releaseTree(root);
}