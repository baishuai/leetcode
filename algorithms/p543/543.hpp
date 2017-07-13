
#ifndef LEETCODE_543_HPP
#define LEETCODE_543_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>
#include "../common/leetcode.hpp"

using namespace std;

/**
Given a binary tree, you need to compute the length of the diameter of the tree.
 The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
 This path may or may not pass through the root.

Example:
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.

 */

class Solution {
public:
    int diameterOfBinaryTree(TreeNode *root) {
        return depthAndDia(root).second;
    }

private:
    /*
     * return <depth, diameterLength>
     */
    pair<int, int> depthAndDia(TreeNode *root) {
        if (root == nullptr) {
            return {0, 0};
        }
        auto l = depthAndDia(root->left);
        auto r = depthAndDia(root->right);
        return {1 + max(l.first, r.first), max(l.first + r.first, max(l.second, r.second))};
    }
};


#endif //LEETCODE_543_HPP
