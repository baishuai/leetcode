//
// Created by bai on 17-6-28.
//

#include "437.hpp"
#include <gtest/gtest.h>

TEST(p437, example) {
    Solution s;

    TreeNode *root = new TreeNode(1);
    root->left = new TreeNode(1);
    root->left->left = new TreeNode(1);

    ASSERT_EQ(2, s.pathSum(root, 2));
    releaseTree(root);
}
