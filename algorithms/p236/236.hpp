//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_236_HPP
#define LEETCODE_236_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

#include "../common/leetcode.hpp"

using namespace std;

class Solution {
public:
    TreeNode *lowestCommonAncestor(TreeNode *root, TreeNode *p, TreeNode *q) {
        if (root == nullptr || root == p || root == q) {
            return root;
        }
        auto left = lowestCommonAncestor(root->left, p, q);
        auto right = lowestCommonAncestor(root->right, p, q);
        return left == nullptr ? right : right == nullptr ? left : root;
    }
};



#endif //LEETCODE_236_HPP
