//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_513_HPP
#define LEETCODE_513_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

#include "../common/leetcode.hpp"

using namespace std;


class Solution {
private:
    int depth;
    int val;

    void dfsBlValue(TreeNode *node, int d) {
        if (node == nullptr) {
            return;
        }
        if (d > depth) {
            depth = d;
            val = node->val;
        }
        dfsBlValue(node->left, d + 1);
        dfsBlValue(node->right, d + 1);
    }

public:
    Solution() : depth(0), val(0) {}

    int findBottomLeftValue(TreeNode *root) {
        depth = 0;
        dfsBlValue(root, 1);
        return val;
    }
};


#endif //LEETCODE_513_HPP
