
#ifndef LEETCODE_230_HPP
#define LEETCODE_230_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>
#include <functional>
#include "../common/leetcode.hpp"

using namespace std;

/**
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note:
You may assume k is always valid, 1 ? k ? BST's total elements.
 */

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

class Solution {
public:
    int kthSmallest(TreeNode *root, int k) {
        int res = 0;
        preOrderDfs(root, [&](int v) -> bool {
            if (--k == 0) {
                res = v;
                return true;
            }
            return false;
        });
        return res;
    }

private:
    void preOrderDfs(TreeNode *root, function<bool(int)> f) {
        if (root != nullptr) {
            preOrderDfs(root->left, f);
            if (!f(root->val)) {
                preOrderDfs(root->right, f);
            }
        }
    }
};

#endif //LEETCODE_230_HPP
