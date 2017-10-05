
#ifndef LEETCODE_530_HPP
#define LEETCODE_530_HPP

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
    int getMinimumDifference(TreeNode *root) {
        if (root == nullptr)
            return min_diff;
        if (root->left != nullptr)
            getMinimumDifference(root->left);
        if (val >= 0)
            min_diff = min(min_diff, root->val - val);
        val = root->val;
        if (root->right != nullptr)
            getMinimumDifference(root->right);
        return min_diff;
    }

private:
    int min_diff{numeric_limits<int>::max()}, val{-1};
};


#endif //LEETCODE_530_HPP
