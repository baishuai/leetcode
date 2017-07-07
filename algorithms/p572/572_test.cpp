//
// Created by bai on 17-6-28.
//

#include "572.hpp"
#include <gtest/gtest.h>

TEST(p572, example) {
    Solution s;

    TreeNode *root1 = new TreeNode(1);
    root1->left = new TreeNode(2);

    TreeNode *root2 = new TreeNode(2);

    ASSERT_TRUE(s.isSubtree(root1, root2));

    releaseTree(root1);
    releaseTree(root2);

}