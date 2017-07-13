
#include "543.hpp"
#include <gtest/gtest.h>

TEST(p543, example) {
    Solution s;
    TreeNode *root = new TreeNode(1);
    root->right = new TreeNode(3);
    root->left = new TreeNode(2);
    root->left->left = new TreeNode(4);
    root->left->right = new TreeNode(5);
    ASSERT_EQ(3, s.diameterOfBinaryTree(root));
    releaseTree(root);
}
