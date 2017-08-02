
#include "404.hpp"
#include <gtest/gtest.h>

TEST(p404, example) {
    Solution s;
    TreeNode *root = new TreeNode(1);
    root->left = new TreeNode(2);
    ASSERT_EQ(2, s.sumOfLeftLeaves(root));
    releaseTree(root);
}
