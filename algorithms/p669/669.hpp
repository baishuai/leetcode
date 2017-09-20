
#ifndef LEETCODE_669_HPP
#define LEETCODE_669_HPP

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
    TreeNode *trimBST(TreeNode *root, int L, int R) {
        if (root == nullptr)
            return nullptr;
        if (root->val < L)
            return trimBST(root->right, L, R);
        if (root->val > R)
            return trimBST(root->left, L, R);
        root->left = trimBST(root->left, L, R);
        root->right = trimBST(root->right, L, R);
        return root;
    }
};

#endif //LEETCODE_669_HPP
