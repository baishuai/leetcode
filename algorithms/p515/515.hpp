
#ifndef LEETCODE_515_HPP
#define LEETCODE_515_HPP

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
You need to find the largest value in each row of a binary tree.

Example:
Input:

          1
         / \
        3   2
       / \   \
      5   3   9

Output: [1, 3, 9]
 */


class Solution {
public:
    vector<int> largestValues(TreeNode *root) {
        vector<int> result;
        dfs(root, result, 0);
        return result;
    }

private:
    void dfs(TreeNode *node, vector<int> &lvr, int depth) {
        if (node == nullptr) {
            return;
        }
        if (lvr.size() == depth) {
            lvr.emplace_back(node->val);
        } else {
            lvr[depth] = max(lvr[depth], node->val);
        }
        dfs(node->left, lvr, depth + 1);
        dfs(node->right, lvr, depth + 1);
    }
};

#endif //LEETCODE_515_HPP
