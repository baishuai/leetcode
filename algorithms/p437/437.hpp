//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_437_HPP
#define LEETCODE_437_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

#include "../common/leetcode.hpp"

using namespace std;

class Solution {

public:
    int pathSum(TreeNode *root, int sum) {
        psums.clear();
        psums[0]++;
        return pathSumR(root, sum, 0);
    }

private:
    unordered_map<int, int> psums;

    int pathSumR(TreeNode *node, int sum, int ps) {
        if (node == nullptr) {
            return 0;
        }
        ps += node->val;

        int cnt = psums[ps - sum];
        psums[ps]++;
        cnt += pathSumR(node->left, sum, ps);
        cnt += pathSumR(node->right, sum, ps);
        psums[ps]--;
        return cnt;
    }
};

#endif //LEETCODE_437_HPP
