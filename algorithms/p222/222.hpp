//
// Created by bai on 17-6-27.
//

#ifndef LEETCODE_222_HPP
#define LEETCODE_222_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <cmath>

#include "../common/leetcode.hpp"

using namespace std;

class Solution {
private:
    int leftDepth(TreeNode *node) {
        if (node == nullptr) {
            return 0;
        }
        int depth = 1;
        while (node->left != nullptr) {
            depth++;
            node = node->left;
        }
        return depth;
    }

    int rightDepth(TreeNode *node) {
        if (node == nullptr) {
            return 0;
        }
        int depth = 1;
        while (node->right != nullptr) {
            depth++;
            node = node->right;
        }
        return depth;
    }

    int leftCNodes(TreeNode *left, int ld) {
        int rd = rightDepth(left);
        if (rd == ld) {
            return (1 << (unsigned) ld) - 1;
        } else {
            return 1 + leftCNodes(left->left, ld - 1) + countNodes(left->right);
        }
    }

public:
    int countNodes(TreeNode *root) {
        int ld = leftDepth(root);
        int rd = rightDepth(root);
        if (ld == rd) {
            return (1 << (unsigned) ld) - 1;
        } else {
            return 1 + leftCNodes(root->left, ld - 1) + countNodes(root->right);
        }
    }
};

#endif //LEETCODE_222_HPP
