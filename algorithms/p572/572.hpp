//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_572_HPP
#define LEETCODE_572_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

#include "../common/leetcode.hpp"

using namespace std;


class Solution {
public:
    bool isSubtree(TreeNode *s, TreeNode *t) {
        if (s == nullptr) {
            return t == nullptr;
        }
        return this->eqTree(s, t) || isSubtree(s->left, t) || isSubtree(s->right, t);
    }

private:
    bool eqTree(TreeNode *lhs, TreeNode *rhs) {
        if (lhs == nullptr && rhs == nullptr) {
            return true;
        } else if (lhs == nullptr || rhs == nullptr) {
            return false;
        }
        return lhs->val == rhs->val && this->eqTree(lhs->left, rhs->left) && this->eqTree(lhs->right, rhs->right);
    }
};

#endif //LEETCODE_572_HPP
