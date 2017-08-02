
#ifndef LEETCODE_404_HPP
#define LEETCODE_404_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>
#include <stack>
#include <string>
#include "../common/leetcode.hpp"

using namespace std;

/*
Find the sum of all left leaves in a given binary tree.

Example:

    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
 */

class Solution {
public:
    int sumOfLeftLeaves(TreeNode *root) {
        return dfsLeftLeaves(root, false);
    }

private:
    int dfsLeftLeaves(TreeNode *node, bool left) {
        if (node == nullptr)
            return 0;
        else if (node->left == nullptr && node->right == nullptr) {
            return left ? node->val : 0;
        } else {
            return dfsLeftLeaves(node->left, true) + dfsLeftLeaves(node->right, false);
        }
    }
};

#endif //LEETCODE_404_HPP
