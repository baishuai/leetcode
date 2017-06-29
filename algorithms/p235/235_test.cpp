//
// Created by bai on 17-6-28.
//

#include "235.hpp"
#include <gtest/gtest.h>

TEST(p235, example) {
    Solution s;

    TreeNode *root = new TreeNode(6);
    TreeNode *l1 = new TreeNode(2);
    TreeNode *r1 = new TreeNode(8);
    root->left = l1;
    root->right = r1;
    TreeNode *lc = new TreeNode(4);
    l1->right = lc;

    ASSERT_EQ(root, s.lowestCommonAncestor(root, root, l1));
    ASSERT_EQ(root, s.lowestCommonAncestor(root, root, r1));
    ASSERT_EQ(root, s.lowestCommonAncestor(root, l1, r1));
    ASSERT_EQ(root, s.lowestCommonAncestor(root, lc, r1));


    ASSERT_EQ(l1, s.lowestCommonAncestor(root, l1, lc));

    releaseTree(root);
}