
#include "230.hpp"
#include <gtest/gtest.h>

TEST(p230, example) {
    Solution s;
    TreeNode *root = new TreeNode(1);
    ASSERT_EQ(1, s.kthSmallest(root, 1));
    releaseTree(root);
}
