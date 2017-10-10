
#ifndef LEETCODE_655_HPP
#define LEETCODE_655_HPP

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


class Solution {
public:
    vector<vector<string>> printTree(TreeNode *root) {
        vector<vector<string>> res;
        if (root == nullptr)
            return res;
        int height = getHeight(root);
        int width = (1 << height) - 1;
        populateRes(res, 0, width / 2, width / 2, root);

        for (int i = 0; i < height; ++i) {
            while (res[i].size() < width)
                res[i].emplace_back("");
        }
        return res;
    }

    void populateRes(vector<vector<string>> &res, int depth, int leftPadding, int levelPadding, TreeNode *node) {
        if (res.size() == depth)
            res.emplace_back(vector<string>{});
        auto &line = res[depth];
        while (line.size() < leftPadding)
            line.emplace_back("");
        line.emplace_back(to_string(node->val));
        if (node->left != nullptr)
            populateRes(res, depth + 1, leftPadding - levelPadding / 2 - 1, levelPadding / 2, node->left);
        if (node->right != nullptr)
            populateRes(res, depth + 1, leftPadding + levelPadding / 2 + 1, levelPadding / 2, node->right);
    }

    int getHeight(TreeNode *node) {
        if (node == nullptr)
            return 0;
        return 1 + max(getHeight(node->left), getHeight(node->right));
    }
};

#endif //LEETCODE_655_HPP
