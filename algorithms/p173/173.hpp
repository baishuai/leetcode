
#ifndef LEETCODE_173_HPP
#define LEETCODE_173_HPP

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
    BSTIterator(TreeNode *root) {
        TreeNode *node = root;
        while (node != nullptr) {
            path.push(node);
            node = node->left;
        }
    }

    /** @return whether we have a next smallest number */
    bool hasNext() {
        return !path.empty();
    }

    /** @return the next smallest number */
    int next() {
        auto t = path.top();
        path.pop();
        if (t->right != nullptr) {
            path.push(t->right);
            auto n = t->right->left;
            while (n != nullptr) {
                path.push(n);
                n = n->left;
            }
        }
        return t->val;
    }

private:
    stack<TreeNode *> path;
};

#endif //LEETCODE_173_HPP
