
#ifndef LEETCODE_257_HPP
#define LEETCODE_257_HPP

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
    vector<string> binaryTreePaths(TreeNode *root) {
        path.clear();
        vector<string> paths;
        dfs(paths, root);
        return paths;
    }

private:
    deque<int> path;

    string path2s() {
        string s;
        for (int i = 0; i < path.size(); ++i) {
            s.append(to_string(path[i]));
            s.append("->");
        }
        return s;
    }

    void dfs(vector<string> &paths, TreeNode *node) {
        if (node == nullptr)
            return;
        else if (node->right == nullptr && node->left == nullptr)
            paths.emplace_back(path2s().append(to_string(node->val)));
        else {
            path.emplace_back(node->val);
            dfs(paths, node->left);
            dfs(paths, node->right);
            path.pop_back();
        }
    }
};

#endif //LEETCODE_257_HPP
