
#include "257.hpp"
#include <gtest/gtest.h>

TEST(p257, example) {
    Solution s;
    TreeNode *root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    vector<string> exp = {"1->2", "1->3"};
    ASSERT_EQ(exp, s.binaryTreePaths(root));
    releaseTree(root);
}
