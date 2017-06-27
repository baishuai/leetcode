//
// Created by bai on 17-6-26.
//

#include "226.hpp"
#include <gtest/gtest.h>

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
