
#include "501.hpp"
#include <gtest/gtest.h>

TEST(p501, example) {
    Solution s;

    TreeNode *root = new TreeNode(1);
    root->right = new TreeNode(2);
    root->right->left = new TreeNode(2);

    vector<int> exp = {2};
    ASSERT_EQ(exp, s.findMode(root));
    releaseTree(root);
}
