
#ifndef LEETCODE_508_HPP
#define LEETCODE_508_HPP

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
Given the root of a tree, you are asked to find the most frequent subtree sum.
 The subtree sum of a node is defined as the sum of all the node values formed
 by the subtree rooted at that node (including the node itself).
 So what is the most frequent subtree sum value? If there is a tie,
 return all the values with the highest frequency in any order.

Examples 1
Input:

  5
 /  \
2   -3
return [2, -3, 4], since all the values happen only once,
 return all of them in any order.
Examples 2
Input:

  5
 /  \
2   -5
return [2], since 2 happens twice, however -5 only occur once.
Note: You may assume the sum of values in any subtree is
 in the range of 32-bit signed integer.
 */

class Solution {
public:
    vector<int> findFrequentTreeSum(TreeNode *root) {
        sumCnt.clear();
        sum(root);
        int maxSum = numeric_limits<int>::min();
        for (auto &p : sumCnt) {
            maxSum = max(maxSum, p.second);
        }
        vector<int> sums;
        for (auto &p : sumCnt) {
            if (p.second == maxSum) {
                sums.emplace_back(p.first);
            }
        }
        return sums;
    }

private:
    unordered_map<int, int> sumCnt;

    int sum(TreeNode *root) {
        if (root == nullptr) {
            return 0;
        }
        int l = sum(root->left), r = sum(root->right);
        sumCnt[l + r + root->val]++;
        return l + r + root->val;
    }
};

#endif //LEETCODE_508_HPP
