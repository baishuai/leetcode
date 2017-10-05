
#ifndef LEETCODE_654_HPP
#define LEETCODE_654_HPP

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
Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

The root is the maximum number in the array.
The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
Construct the maximum tree by the given array and output the root node of this tree.

Example 1:
Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:

      6
    /   \
   3     5
    \    /
     2  0
       \
        1
Note:
The size of the given array will be in the range [1,1000].
 */


class Solution {
public:
    TreeNode *constructMaximumBinaryTree(vector<int> &nums) {
        stack<TreeNode *> stk;
        for (int i = 0; i < nums.size(); ++i) {
            TreeNode *cur = new TreeNode(nums[i]);
            while (!stk.empty() && stk.top()->val < nums[i]) {
                cur->left = stk.top();
                stk.pop();
            }
            if (!stk.empty())
                stk.top()->right = cur;
            stk.push(cur);
        }
        TreeNode *root = stk.top();
        while (!stk.empty()) {
            root = stk.top();
            stk.pop();
        }
        return root;
    }
};


#endif //LEETCODE_654_HPP
