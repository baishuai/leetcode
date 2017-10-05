
#ifndef LEETCODE_563_HPP
#define LEETCODE_563_HPP

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
Given a binary tree, return the tilt of the whole tree.

The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.

The tilt of the whole tree is defined as the sum of all nodes' tilt.

Example:
Input:
         1
       /   \
      2     3
Output: 1
Explanation:
Tilt of node 2 : 0
Tilt of node 3 : 0
Tilt of node 1 : |2-3| = 1
Tilt of binary tree : 0 + 0 + 1 = 1
Note:

The sum of node values in any subtree won't exceed the range of 32-bit integer.
All the tilt values won't exceed the range of 32-bit integer.
 */


class Solution {
public:
    int findTilt(TreeNode *root) {
        tilt = 0;
        dfsSum(root);
        return tilt;
    }

private:
    int tilt;

    int dfsSum(TreeNode *node) {
        if (node == nullptr)
            return 0;
        int l = dfsSum(node->left);
        int r = dfsSum(node->right);
        tilt += abs(l - r);
        return l + r + node->val;
    }
};

#endif //LEETCODE_563_HPP
