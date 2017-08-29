
#ifndef LEETCODE_662_HPP
#define LEETCODE_662_HPP

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
 Given a binary tree, write a function to get the maximum width of the given tree.
 The width of a tree is the maximum width among all levels.
 The binary tree has the same structure as a full binary tree, but some nodes are null.

The width of one level is defined as the length between the end-nodes
 (the leftmost and right most non-null nodes in the level,
 where the null nodes between the end-nodes are also counted into the length calculation.
 */


class Solution {
public:
    int widthOfBinaryTree(TreeNode *root) {
        if (root == nullptr)
            return 0;

        queue<int> count;
        queue<TreeNode *> level;
        level.push(root);
        count.push(0);
        int width = 1;

        while (!level.empty()) {
            auto size = level.size();
            int left = 0;
            int right = 0;
            for (int i = 0; i < size; ++i) {
                auto node = level.front();
                int index = count.front();
                level.pop();
                count.pop();
                if (i == 0)
                    left = index;
                if (i == size - 1)
                    right = index;
                if (node->left != nullptr) {
                    level.push(node->left);
                    count.push(index * 2);
                }
                if (node->right != nullptr) {
                    level.push(node->right);
                    count.push(index * 2 + 1);
                }
            }
            width = max(width, right - left + 1);
        }
        return width;
    }

};


#endif //LEETCODE_662_HPP
