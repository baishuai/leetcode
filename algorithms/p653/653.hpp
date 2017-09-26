
#ifndef LEETCODE_653_HPP
#define LEETCODE_653_HPP

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

class BSTIterator {
public:
    BSTIterator(TreeNode *node, bool forward) : node(node), forward(forward) { (*this)++; }

    int operator*() {
        return cur;
    }

    void operator++(int) {
        while (node || !depth.empty()) {
            if (node) {
                depth.emplace(node);
                node = forward ? node->left : node->right;
            } else {
                node = depth.top();
                depth.pop();
                cur = node->val;
                node = forward ? node->right : node->left;
                break;
            }
        }
    }

private:
    stack<TreeNode *> depth;
    TreeNode *node;
    bool forward;
    int cur;
};

class Solution {
public:
    bool findTarget(TreeNode *root, int k) {
        if (root == nullptr)
            return false;
        BSTIterator l(root, true);
        BSTIterator r(root, false);
        while (*l < *r) {
            if (*l + *r == k)
                return true;
            else if (*l + *r < k)
                l++;
            else
                r++;
        }
        return false;
    }
};

#endif //LEETCODE_653_HPP
