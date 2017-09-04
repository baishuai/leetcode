
#ifndef LEETCODE_637_HPP
#define LEETCODE_637_HPP

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
Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.

Example 1:
Input:
    3
   / \
  9  20
    /  \
   15   7
Output: [3, 14.5, 11]
Explanation:
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11.
 Hence return [3, 14.5, 11].
 */

class Solution {
public:
    vector<double> averageOfLevels(TreeNode *root) {
        num.clear();
        sum.clear();
        dfs(root, 0);
        vector<double> res(num.size());
        for (int i = 0; i < num.size(); ++i)
            res[i] = static_cast<double >(sum[i]) / static_cast<double >(num[i]);
        return res;
    }

private:
    vector<int> num;
    vector<long long> sum;

    void dfs(TreeNode *node, int depth) {
        if (node == nullptr)
            return;
        if (num.size() == depth) {
            num.emplace_back(1);
            sum.emplace_back(node->val);
        } else {
            num[depth]++;
            sum[depth] += node->val;
        }
        dfs(node->left, depth + 1);
        dfs(node->right, depth + 1);
    }
};

#endif //LEETCODE_637_HPP
