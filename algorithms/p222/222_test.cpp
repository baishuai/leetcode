//
// Created by bai on 17-6-27.
//

#include "222.hpp"
#include <gtest/gtest.h>

TEST(p222, example) {
    Solution s;

    TreeNode *root = new TreeNode(1);
    root->left = new TreeNode(2);
    ASSERT_EQ(2, s.countNodes(root));
    releaseTree(root);
}