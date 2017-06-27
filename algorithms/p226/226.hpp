//
// Created by bai on 17-6-26.
//

#ifndef LEETCODE_226_HPP
#define LEETCODE_226_HPP

/**
 * Invert a binary tree.

     4
   /   \
  2     7
 / \   / \
1   3 6   9
to
     4
   /   \
  7     2
 / \   / \
9   6 3   1

 Trivia:

 This problem was inspired by this original tweet by Max Howell:
Google: 90% of our engineers use the software you wrote (Homebrew),
but you can’t invert a binary tree on a whiteboard so fuck off.

 */


#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

#include "../common/leetcode.hpp"

using namespace std;

class Solution {
public:
    TreeNode *invertTree(TreeNode *root) {
        if (root != nullptr) {
            TreeNode *tmp = root->left;
            root->left = invertTree(root->right);
            root->right = invertTree(tmp);
        }
        return root;
    }
};

#endif //LEETCODE_226_HPP
