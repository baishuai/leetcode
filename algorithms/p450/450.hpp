//
// Created by bai on 17-6-27.
//

#ifndef LEETCODE_450_HPP
#define LEETCODE_450_HPP

/**
 * Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.
Note: Time complexity should be O(height of tree).

Example:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

Given key to delete is 3. So we find the node with value 3 and delete it.

One valid answer is [5,4,6,2,null,null,7], shown in the following BST.

    5
   / \
  4   6
 /     \
2       7

Another valid answer is [5,2,6,null,4,null,7].

    5
   / \
  2   6
   \   \
    4   7
 */

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

using namespace std;

#include "../common/leetcode.hpp"

class Solution {

    pair<TreeNode *, TreeNode *> findNextChild(TreeNode *node) {
        TreeNode *p = node;
        TreeNode *c = node->right;
        while (c->left != nullptr) {
            p = c;
            c = c->left;
        }
        return pair<TreeNode *, TreeNode *>(p, c);
    }

public:
    TreeNode *deleteNode(TreeNode *root, int key) {
        if (root != nullptr) {
            root->left = deleteNode(root->left, key);
            root->right = deleteNode(root->right, key);
            if (root->val == key) {
                if (root->right == nullptr) {
                    TreeNode *tmp = root;
                    root = root->left;
                    delete tmp;
                } else {
                    auto p = findNextChild(root);
                    if (p.first->left == p.second) {
                        p.first->left = p.second->right;
                    } else if (p.first->right == p.second) {
                        p.first->right = p.second->right;
                    }
                    root->val = p.second->val;
                    delete p.second;
                }
            }
        }
        return root;
    }
};

#endif //LEETCODE_450_HPP
