
#include "508.hpp"
#include <gtest/gtest.h>

TEST(p508, example) {
    Solution s;
    TreeNode *root = new TreeNode(5);
    root->left = new TreeNode(2);
    root->right = new TreeNode(-3);
    vector<int> exp = {-3, 2, 4};

    vector<int> res = s.findFrequentTreeSum(root);
    sort(res.begin(), res.end());
    ASSERT_EQ(exp, res);
    releaseTree(root);
}
