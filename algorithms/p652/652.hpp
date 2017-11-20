
#ifndef LEETCODE_652_HPP
#define LEETCODE_652_HPP

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
    vector<TreeNode *> findDuplicateSubtrees(TreeNode *root) {
        trees.clear();
        res.clear();
        postorder(root);
        return res;
    }

private:
    unordered_map<string, int> trees;
    vector<TreeNode *> res;

    string postorder(TreeNode *node) {
        if (node == nullptr) {
            return "#";
        }
        string serial = to_string(node->val) + "," + (postorder(node->left)) + "," + (postorder(node->right));
        trees[serial]++;
        if (trees[serial] == 2)
            res.emplace_back(node);
        return serial;
    }
};

#endif //LEETCODE_652_HPP
